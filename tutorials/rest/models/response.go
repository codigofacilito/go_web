package models

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Response struct{
	Status 	    int 		       `json:"status"`
	Data 	      interface{}	   `json:"data"`
	Message     string		     `json:"message"`
  contentType string
  writer      http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response{
	return Response{ Status: http.StatusOK, writer: w, contentType: "application/json" }
}

func SendNotFound(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.NotFound()
  response.Send()
}

func (this *Response) NotFound(){
  this.Status = http.StatusNotFound
  this.Message = "Resource Not Found."
}

func SendUnprocessableEntity(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.UnprocessableEntity()
  response.Send()
}

func (this *Response) UnprocessableEntity(){
  this.Status = http.StatusUnprocessableEntity
  this.Message = "Unprocessable Entity."
}

func SendNoContent(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.NoContent()
  response.Send()
}

func (this *Response) NoContent(){
  this.Status = http.StatusNoContent
  this.Message = "No Content."
} 

func SendData(w http.ResponseWriter, data interface{}){
  response := CreateDefaultResponse(w)
  response.Data = data
  response.Send()
}

func (this *Response) Send(){
  this.writer.Header().Set("Content-Type", this.contentType)
  this.writer.WriteHeader(this.Status)

  output, _ := json.Marshal(&this)
  fmt.Fprintf(this.writer, string(output) )
}

