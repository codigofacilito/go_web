package handlers

import(
  "fmt"
  "../utils"
  "net/http"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(function customeHandler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    
    if !utils.IsAuthenticated(r){
      http.Redirect(w,r, "/users/login", http.StatusSeeOther)
      return
    }
    function(w, r)
  })
}

func MiddlewateTwo(handler http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    //logia aquí!
    fmt.Println("Este es el segundo wrap!")
    handler.ServeHTTP(w, r)
  });
}
