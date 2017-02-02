package main

import(
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request){

		r.Header.Set("name", "value")
		fmt.Println(r.Header) 

		accept := r.Header.Get("Accept")
		if len(accept) != 0{
			fmt.Println(accept)
		}
		fmt.Fprintf(w, "Response")
	})
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}
