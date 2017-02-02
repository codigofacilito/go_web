package main

import(
    "net/http"
    "html/template"
    "log"
)

type Course struct{
    Name string
}

type User struct{
    UserName string
    Activate bool
    Tags []string
    Courses []Course
}

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        template, error := template.ParseFiles("index.html")
        if error != nil{
            panic(error)
        }

        tags := []string{"Go", "Web", "CodigoFacilito"}
        cursos := []Course{
                Course{"Python"}, Course{"Go"}, Course{"Flask"},
        }
        user := User{
            UserName: "Eduardo Ismael",
            Activate : true,
            Courses: cursos,
            Tags: tags,
        }
        template.Execute(w, user)
    })
        
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
