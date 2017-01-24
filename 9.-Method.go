package main

import(
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("El método es : " + r.Method)
		switch r.Method {
			case "GET":
			    fmt.Fprintf(w, "Hola Mundo con el método get")
			case "POST":
			    fmt.Fprintf(w, "Hola Mundo con el método Post")
			case "PUT":
			    fmt.Fprintf(w, "Hola Mundo con el método Put")
			case "DELETE":
			    fmt.Fprintf(w, "Hola Mundo con el método Delete")
			default:
				http.Error(w, "Método no valido.", http.StatusBadRequest)
		}
	})
		
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}