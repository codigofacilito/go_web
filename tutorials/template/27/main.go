package main

import(
	"net/http"
    "html/template"
	"log"
)

type User struct{
    UserName string
}

func (this *User) userNameFormat() string{
    return "El nombre del usuarios es "+ this.UserName;
}

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        template, error := template.ParseFiles("index.html")
        if error != nil{
            panic(error)
        }
            
        user := User{"Eduardo Ismael."}
        template.Execute(w, user)
    })
    	
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
