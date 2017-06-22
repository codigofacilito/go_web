package models

import "errors"

type ValidationError error

var(
  errorUsername = ValidationError(errors.New("El username no debe de esta vacío.") )
  errorShortUsername = ValidationError(errors.New("El username es demasiado corto.") )
  errorLargeUsername = ValidationError(errors.New("El username es demasiado largo.") )
  
  errorEmail = ValidationError(errors.New("Formato invalido de Email."))

  errorPasswordEncryption = ValidationError(errors.New("No es posible cifrar el texto."))
)