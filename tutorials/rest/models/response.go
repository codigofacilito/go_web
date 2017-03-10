package models

import "net/http"

type Response struct{
	Status 	int 		`json:"status"`
	Data 	interface{}	`json:"data"`
	Message string		`json:"message"`
}

func GetDefaultResponse() Response{
	return Response{ Status: http.StatusOK }
}


func (this *Response) NotFound(message string){
	this.Status = http.StatusNotFound
	this.Data = nil
	this.Message = message
}