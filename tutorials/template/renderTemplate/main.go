package main

import(
    "net/http"
    "html/template"
    "log"
)

type Usuario struct{
    Username string
}

var templates = template.Must( template.New("T").ParseGlob("templates/**/*.html") )
var error = template.Must( template.ParseFiles("templates/error.html") )

func handlerError(w http.ResponseWriter, status int){
    w.WriteHeader(status)
    error.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}){

    w.Header().Set("Content-Type", "text/html")
    err := templates.ExecuteTemplate(w, name , data)        
    if err != nil{
        handlerError(w, http.StatusInternalServerError)
    }
}

func main(){

    http.HandleFunc("/usuario", func(w http.ResponseWriter, r * http.Request){
        usuario := Usuario{ "Eduardo" }
        renderTemplate(w, "usuario", usuario)
    })
            
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
