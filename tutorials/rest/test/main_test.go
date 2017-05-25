package test

import(
  "os"
  "fmt"
  "testing"
  "../models"
)

func TestMain(m *testing.M){
  beforeTest()
  result := m.Run()
  afterTest()
  os.Exit(result)
}

func beforeTest(){
  fmt.Println(">> Antes de las pruebas")
  models.CreateConnection()
  models.CreateTables()
  createDefaultUser()
}

func createDefaultUser(){
  sql := fmt.Sprintf("INSERT users SET id='%d', username='%s', password='%s', email='%s', created_date='%s' ", id, username, passwordHash, email, createdDate)
  _, err := models.Exec(sql)
  if err != nil{
    panic(err)
  }
  user = &models.User{ Id:id, Username:username, Password:password, Email:email }
}

func afterTest(){
  fmt.Println(">> Después de las pruebas")
  models.CloseConnection()
}



