package validation

import (
	"errors"
	"fmt"
	"io/ioutil"
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

type Translator struct {
	Bundle *i18n.Bundle
}

func NewTranslator() *Translator {
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

func loadFiles(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to download the file, status code:", resp.StatusCode)
	}

	fileBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the file:", err)
	}

	return fileBytes
}
