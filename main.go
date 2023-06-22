package main

import (
	"fmt"
	"os"

	"github.com/xbizzybone/go-toolkit/validation"
)

type User struct {
	Email string `validate:"required,email" json:"email_json"`
	Name  string `validate:"required" json:"name_json"`
}

func main() {
	translator := validation.NewValidatorMessageTranslator()

	user := User{
		Email: "bad_email",
		Name:  "hola",
	}

	localesEsBytes, _ := os.ReadFile("./validation/locales/active.es.toml")

	translator.AddCustomMustParseMessageFileBytes(localesEsBytes, "active.es.toml")
	translator.AddCustomMustParseMessageFileBytes(localesEsBytes, "active.en.toml")

	fmt.Println(translator.ValidateSchema("jp", user)) // not supported language
	fmt.Println(translator.ValidateSchema("en", user))
	fmt.Println(translator.ValidateSchema("es", user))
}
