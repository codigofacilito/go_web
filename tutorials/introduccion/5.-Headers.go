package main

import(
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("custom", "value")
		fmt.Fprintf(w, "Hola Mundo.")
	})
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}