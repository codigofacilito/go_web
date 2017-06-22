package main

import(
	"log"
	"net/http"
	"./config"
	"./handlers"
	"./handlers/api/v1"
	"github.com/gorilla/mux"
)

func main() {
	
	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.Index).Methods("GET")
	mux.HandleFunc("/users/new", handlers.NewUser).Methods("GET", "POST")

	mux.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", v1.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")
	
	log.Println("El servidor esta a la escucha en el puerto ", config.ServerPort() )
	server := &http.Server{
		Addr: config.UrlServer(),
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}