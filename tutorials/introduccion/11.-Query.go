package main

import(
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request){
		
		values := r.URL.Query()
		values.Add("course", "Golang")
		values.Del("name")
		
		r.URL.RawQuery = values.Encode()

		fmt.Println(r.URL.RawQuery) //encoded query values, without '?'
		fmt.Println(r.URL)

		fmt.Fprintf(w, "Response")
	})
	
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}