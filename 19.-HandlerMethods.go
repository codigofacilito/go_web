package main

import(
	"net/http"
	"log"
)

func main() {
  mux := http.NewServeMux()
  redirectHandler := http.RedirectHandler("http://codigofacilito.com", http.StatusTemporaryRedirect )
  notFoundHandler := http.NotFoundHandler()
    
  mux.Handle("/re", redirectHandler)
  mux.Handle("/no", notFoundHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

