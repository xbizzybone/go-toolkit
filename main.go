package main

import (
	"fmt"

	"github.com/xbizzybone/go-toolkit/errors"
	"github.com/xbizzybone/go-toolkit/validation"
)

type User struct {
	Email  string  `validate:"required,email" json:"email_json"`
	Name   string  `validate:"required,min=6" json:"name_json"`
	Amount float64 `validate:"required,min=6" json:"amount_json"`
}

func main() {
	translator := validation.NewValidatorMessageTranslator(validation.Struct)
	logger := errors.NewLogger("debug")

	user := User{
		Email:  "email@gmail.com",
		Name:   "hola",
		Amount: 1023,
	}

	fmt.Println(translator.ValidateSchema("jp", user)) // not supported language
	fmt.Println(translator.ValidateSchema("en", user))
	fmt.Println(translator.ValidateSchema("es", user))

	logger.Debug("Debug message", map[string]interface{}{"key": "value"})
	logger.Info("Info message")
	logger.Warn("Warn message")
	logger.Error("Error message", fmt.Errorf("Error"))
}
