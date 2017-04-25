package main

import(
  "fmt"
  "./orm"
)

func main() {
  orm.CreateConnection()
  orm.CreateTables()

  user := orm.CreateUser("eduardo", "123", "eduardo@codigofacilito.com")
  user = orm.GetUser(1)
  fmt.Println(user)

  user.Username = "Codigo Facilito"
  user.Password = "change_password"
  user.Save()
  
  fmt.Println(user)

  user.Delete()

  orm.CloseConnection()

}