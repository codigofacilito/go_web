package main

import(
    "net/http"
    "html/template"
    "log"
)

type User struct{
    UserName string
    Activate bool
    Admin bool
    Edad int
}

func (this *User) TienePermiso(feature string) bool {  
  return this.Activate && this.Admin && feature == "si" ;
}

func hola() string{
    return "Hola desde una función"
}

func suma(v1, v2 int) int{
    return v1 + v2
}

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){

        funciones := template.FuncMap{
            "hola" : hola,
            "suma" : suma,
        }

        //template, error := template.ParseFiles("templates/index.html") //Función
        template, error := template.New("index.html").Funcs(funciones).ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html") //Método

        if error != nil{
            panic(error)
        }
        
        user := User{
            UserName: "Eduardo Ismael",
            Activate : true,
            Admin : true,
            Edad: 20,
        }

        template.Execute(w, user)
    })
        
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
