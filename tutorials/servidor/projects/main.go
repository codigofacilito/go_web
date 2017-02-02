package main

import(
	"net/http"
	"log"
	"fmt"
	"./mux"
)

func Hola(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola desde una función anonima")
}

type User struct{
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola Usuario " + this.name )
}

func main() {

	user := &User{ "Eduardo" }
	mux := mux.CreateMux()

	mux.AddFunc("/hola", Hola)
	mux.AddHandle("/usuario", user)

	log.Fatal(http.ListenAndServe("localhost:3000",  mux ))
}

