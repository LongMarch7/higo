package validator

import (
    "gopkg.in/go-playground/validator.v9"
)
var Validate *validator.Validate

func init(){
    Validate = validator.New()
}

const Email =  "required,email"
const ArrayNumeric =  "required,gt=0,dive,required,numeric"
const Alphanum =  "required,alphanum"
const ArrayAlphanum =  "required,gt=0,dive,required,alphanum"
