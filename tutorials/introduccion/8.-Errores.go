package main

import(
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//Documentacón
		//https://golang.org/src/net/http/status.go

		http.Error(w, "Una breve descripción del error.", http.StatusUnauthorized)
	})
		
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}