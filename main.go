package main

import (
	"fmt"

	"github.com/xbizzybone/go-toolkit/validation"
)

type User struct {
	Email string `validate:"required,email" json:"email_json"`
	Name  string `validate:"required,min=6" json:"name_json"`
}

func main() {
	translator := validation.NewValidatorMessageTranslator(validation.Struct)

	user := User{
		Email: "email@gmail.com",
		Name:  "hola",
	}

	fmt.Println(translator.ValidateSchema("jp", user)) // not supported language
	fmt.Println(translator.ValidateSchema("en", user))
	fmt.Println(translator.ValidateSchema("es", user))
}
