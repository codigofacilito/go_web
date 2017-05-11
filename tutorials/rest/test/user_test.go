package test


import(
  "os"
  "testing"
  "../models"
)

func TestMain(m *testing.M){
  setupFuncition()
  result := m.Run()
  teardownFunction()
  os.Exit(result)
}

func setupFuncition(){
  /*Lo que necesitamos antes*/
}

func teardownFunction(){
  /*Lo que necesitamos después*/
}

func TestNewUser(t *testing.T){
  user := models.NewUser("username","password","email")
  if user.Username != "usernames" || user.Password != "password" || user.Email != "email"{
    t.Error("No es posible crear el objeto")
  }
}

func TestSave(t *testing.T){
  user := models.NewUser("username", "password", "email")
  if err := user.Save(); err != nil{
    t.Error("No es posible crear el usuario", err)
  }
}

func TestCreateUser(t *testing.T){
  _, err := models.CreateUser("username", "password", "email")
  if err != nil{
    t.Error("No es posible insertar el objeto", nil)
  }
}