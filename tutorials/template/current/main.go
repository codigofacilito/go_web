package main

import(
    "net/http"
    "html/template"
    "log"
)

func hello() string{
    return "Hola Mundo";
}

func suma(v1, v2 int) int{
    return v1 + v2;
}

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        
        funcs := template.FuncMap{
            "hola": hello,
            "suma": suma,
        }

        template, err := template.New("index.html").Funcs(funcs).ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")
        if err != nil{
            panic(err)
        }
        template.Execute(w, nil)
    })
        
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
