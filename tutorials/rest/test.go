package main

import(
  "./models"
)

func main() {
  models.CreateConnection()
  models.Ping()
  
  models.CreateTables()

  models.CloseConnection()
}