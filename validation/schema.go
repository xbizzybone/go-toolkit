package validation

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var v = validator.New()
var filesURL = "https://raw.githubusercontent.com/xbizzybone/go-toolkit/master/validation/locales"

type Language string

const (
	English Language = "en"
	Spanish Language = "es"
)

type ValidatorMessageTranslator struct {
	Bundle *i18n.Bundle
}

func NewValidatorMessageTranslator() *ValidatorMessageTranslator {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustParseMessageFileBytes(loadFiles(filesURL+"/active.en.toml"), "active.en.toml")
	bundle.MustParseMessageFileBytes(loadFiles(filesURL+"/active.es.toml"), "active.es.toml")

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &ValidatorMessageTranslator{
		Bundle: bundle,
	}
}

func (l Language) isSupported() bool {
	switch l {
	case English, Spanish:
		return true
	default:
		return false
	}
}

func (t *ValidatorMessageTranslator) ValidateSchema(lang string, data interface{}) error {
	if !Language(lang).isSupported() {
		return errors.New("language not supported")
	}

	if err := v.Struct(data); err != nil {
		return errors.New(t.validatorFormatError(lang, err))
	}
	return nil
}

func (t *ValidatorMessageTranslator) validatorFormatError(lang string, err error) string {
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

func loadFiles(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to download the file, status code:", resp.StatusCode)
	}

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the file:", err)
	}

	return fileBytes
}

func isValidFilename(filename string) bool {
	filenameSlice := strings.Split(filename, ".")

	if len(filenameSlice) != 3 {
		return false
	}

	if filenameSlice[0] != "active" {
		return false
	}

	if filenameSlice[2] != "toml" {
		return false
	}

	return true
}

/*
AddCustomMustParseMessageFileBytes is a function that allows you to add custom messages to the validator.
The fileBytes parameter is the content of the file to be added.
The filename parameter is the name of the file to be added. It must be in the format: active.{language}.toml (e.g. active.fr.toml)

The format of the file must be as follows:
[is-hola] # tag
one = "The field %s must be hola" # message
other = "The field %s must be hola" # message

[is-required]
one = "The field %s is required"
other = "The field %s is required"
*/
func (t *ValidatorMessageTranslator) AddCustomMustParseMessageFileBytes(fileBytes []byte, filename string) {
	if !isValidFilename(filename) {
		panic("invalid filename")
	}

	t.Bundle.MustParseMessageFileBytes(fileBytes, filename)
}

/*
AddCustomMustParseMessageFileBytesFromURL is a function that allows you to add custom messages to the validator.
The url parameter is the url of the file to be added.

The format of the file must be as follows:
[is-hola] # tag
one = "The field %s must be hola" # message
other = "The field %s must be hola" # message

[is-required]
one = "The field %s is required"
other = "The field %s is required"
*/
func (t *ValidatorMessageTranslator) AddCustomMustParseMessageFileBytesFromURL(url string, filename string) {
	if !isValidFilename(filename) {
		panic("invalid filename")
	}

	fileBytes := loadFiles(url)
	t.AddCustomMustParseMessageFileBytes(fileBytes, filename)
}
