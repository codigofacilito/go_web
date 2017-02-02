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

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        template, error := template.ParseFiles("index.html")// creates a new Template a
        if error != nil{
            panic(error)
        }
        /*
        and     &&
        or      ||
        eq      ==
        ne      !=
        lt      <
        le      <=
        gt      >
        ge      >=
        */
        user := User{
            UserName: "Eduardo Ismael",
            Activate : true,
            Admin : false,
            Edad: 20,
        }

        template.Execute(w, user)
    })
    	
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
