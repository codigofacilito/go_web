package main

import(
  _ "fmt"
  "./models"
)

func main() {
  models.CreateConnection()
  models.Ping()
}