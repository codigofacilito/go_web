package main

import(
	"net/http"
    "html/template"
	"log"
)

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
        //http://goinbigdata.com/example-of-using-templates-in-golang/
        //https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html
        
        determinar el aspecto visual de nuestro sitio web

        //SI O SI
        //https://www.calhoun.io/an-intro-to-templates-in-go-part-1-of-3/
        
        w.Header().Set("Content-Type", "text/html")
        template, error := template.New("hello").Parse("Hola usuario.")
                
        if error != nil{
            panic(error)
        }

        template.Execute(w, nil)
    })
    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", nil))
}
