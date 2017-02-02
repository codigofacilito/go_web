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

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        template, error := template.ParseFiles("index.html")
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
