package main

import(
	"net/http"
	"log"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola Mundo!")
}

func main() {
   	mux := http.NewServeMux()
   	mux.HandleFunc("/", hello)

   	log.Println("El servidor esta a la escucha del puerto 3000")
	server := &http.Server{
		Addr:	":3000",
		Handler:	mux,
	}
	
	log.Fatal(server.ListenAndServe())
}