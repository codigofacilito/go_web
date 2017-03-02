package main

import(
    "net/http"
    "html/template"
    "log"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html") )
var errorTemplate = template.Must(template.ParseFiles("error.html"))


func handlerError(w http.ResponseWriter, r * http.Request, status int){
    w.WriteHeader(status)
    errorTemplate.Execute(w, r.URL)
}

func renderTemplate(w http.ResponseWriter, r * http.Request, name string, data interface{}){
    w.Header().Set("Content-Type", "text/html")
    err := templates.ExecuteTemplate(w, name, data)
    if err != nil{
        handlerError(w, http.StatusInternalServerError)
    }
}

func index(w http.ResponseWriter, r * http.Request){
    renderTemplate(w, r, "users/indexs", nil)
}

func main(){
    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
                    
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", mux))
}
