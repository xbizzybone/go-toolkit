package main

import (
	"fmt"

	"github.com/xbizzybone/go-toolkit/validation"
)

type User struct {
	Email string `validate:"required,email" json:"email_json"`
	Name  string `validate:"required"`
}

func main() {
	translator := validation.NewValidatorMessageTranslator()

	user := User{
		Email: "bad_email",
		Name:  "John Doe",
	}

	fmt.Println(translator.ValidateSchema("jp", user)) // not supported language
	fmt.Println(translator.ValidateSchema("en", user))
	fmt.Println(translator.ValidateSchema("es", user))
}
