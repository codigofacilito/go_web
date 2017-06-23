package handlers

import (
  "fmt"
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

func Login(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})

  if r.Method == "POST"{
    username := r.FormValue("username")
    password := r.FormValue("password")

    if user, err := models.Login(username, password); err != nil{
      context["Error"] = err.Error()
    }else{
      utils.SetSession(user, w)
      fmt.Println("Estas autenticado.")
    }
  }
  utils.RenderTemplate(w, "users/login", context)
}

func Logout(w http.ResponseWriter, r *http.Request){
  utils.DeleteSession(w)
  http.Redirect(w,r, "/users/login", http.StatusSeeOther)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  user := utils.GetUser(r)
  context["User"] = user
  utils.RenderTemplate(w, "users/edit", context)
}