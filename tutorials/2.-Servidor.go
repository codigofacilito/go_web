package main

import(
	"net/http"
	"log"
)

func main() {
	log.Fatal( http.ListenAndServe("localhost:3000", nil))
}