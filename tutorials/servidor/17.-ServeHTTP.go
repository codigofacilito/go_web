package main

import(
	"net/http"
	"log"
	"fmt"
)

type User struct {
  name string
}

func (th *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello " + th.name )
}

func main() {
  
  username := &User{name: "Your name."}
  log.Println("Listening...")
  http.ListenAndServe(":3000", username)
}

