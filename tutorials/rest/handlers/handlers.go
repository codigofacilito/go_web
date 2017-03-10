package handlers

import(
	"fmt"
	"net/http"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se obtiene todos los usuarios!")
}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	userId, _ :=  strconv.Atoi( vars["id"] )  //string

	response := models.GetDefaultResponse()
	user, err := models.GetUser(userId)

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		response.NotFound(err.Error())
	}else{
		response.Data = user
	}

	output, _ := json.Marshal(&response)
	fmt.Fprintf(w, string(output) )
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se crea un usuario!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se actualiza un usuario!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se elimina un usuario!")
}

