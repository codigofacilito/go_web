package main

import(
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w,r, "/redirect", http.StatusMovedPermanently)		
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w,r, "https://www.google.com.mx", http.StatusMovedPermanently)		
	})
	
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}