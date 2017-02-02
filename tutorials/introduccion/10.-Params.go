package main

import(
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request){
		values := r.URL.Query()

		name := r.URL.Query().Get("name")
		if len(name) != 0 {
		    fmt.Println(name)
		}
		fmt.Fprintf(w, "Response")
	})
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}