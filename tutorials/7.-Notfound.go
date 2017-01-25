package main

import(
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.NotFound(w, r)
	})
		
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}