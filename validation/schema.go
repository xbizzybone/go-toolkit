package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var v = validator.New()

type Language string

const (
	English Language = "en"
	Spanish Language = "es"
)

type Translator struct {
	Bundle *i18n.Bundle
}

func NewTranslator() *Translator {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./validation/locales/active.en.toml")
	bundle.MustLoadMessageFile("./validation/locales/active.es.toml")

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Translator{
		Bundle: bundle,
	}
}

func (l Language) IsSupported() bool {
	switch l {
	case English, Spanish:
		return true
	default:
		return false
	}
}

func (t *Translator) ValidateSchema(lang string, data interface{}) error {
	if !Language(lang).IsSupported() {
		return errors.New("language not supported")
	}

	if err := v.Struct(data); err != nil {
		return errors.New(t.validatorFormatError(lang, err))
	}
	return nil
}

func (t *Translator) validatorFormatError(lang string, err error) string {
	loc := i18n.NewLocalizer(t.Bundle, string(lang))
	message := ""

	for _, err := range err.(validator.ValidationErrors) {
		fieldname := err.Field()
		tag := err.Tag()
		errParam := err.Param()
		errValue := err.Value().(string)

		msg, err := loc.Localize(&i18n.LocalizeConfig{
			MessageID:   tag,
			PluralCount: 1,
		})

		if err != nil {
			panic(err)
		}

		if errParam != "" && errValue == "" {
			message += fmt.Sprintf(msg+" \n", fieldname, errParam)
			continue
		}

		if errParam != "" && errValue != "" {
			message += fmt.Sprintf(msg+"\n", fieldname, errParam, errValue)
			continue
		}

		message += fmt.Sprintf(msg+"\n", fieldname)
	}

	return message
}
