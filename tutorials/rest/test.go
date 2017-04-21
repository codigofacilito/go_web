package main

import(
  "fmt"
  "./models"
)

func main() {
  models.CreateConnection()
  models.CreateTables()

  models.CreateUser("eduardo1", "123", "eduardo@codigofacilito.com")
  models.CreateUser("eduardo2", "123", "eduardo@codigofacilito.com")
  models.CreateUser("eduardo3", "123", "eduardo@codigofacilito.com")

  users := models.GetUsers()
  fmt.Println(users)

  models.CloseConnection()
}