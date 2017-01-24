package main

import(
	"net/http"
	"log"
	"fmt"
	)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Función Hola Mundo.")
}

func main() {
	
   	http.HandleFunc("/", hello)
   	
   	log.Println("El servidor esta a la escucha del puerto 3000")
	server := &http.Server{
		Addr:	"localhost:3000",
		Handler:	nil,
	}
	log.Fatal(server.ListenAndServe())
}