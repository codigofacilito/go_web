package test

import(
  "testing"
)

func TestHolaMundo(t *testing.T){
  str := "hola mundo"
  
  if str != "hola mundo"{ //Esta es la prueba
    t.Error("No es posible saludar a los usuarios", nil)
  }
}