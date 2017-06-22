package handlers

import (
  "../utils"
  "net/http"
  "../models"
)

func NewUser(w http.ResponseWriter, r *http.Request){
  context := make(map[string] interface{})

  if r.Method == "POST"{
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")

    if _, err := models.CreateUser(username, password, email); err != nil{
      errorMessage := err.Error()
      context["Error"] = errorMessage
    }
  }
  utils.RenderTemplate(w, "users/new", context)
}