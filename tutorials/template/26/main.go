package main

import(
	"net/http"
    "html/template"
	"log"
)

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        template, error := template.ParseFiles("index.html")
        if error != nil{
            panic(error)
        }
        template.Execute(w, nil)
    })
    	
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
