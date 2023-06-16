package validation

import (
	"errors"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func validateSchema(data interface{}) error {
	if err := Validator.Struct(data); err != nil {
		log.Println(err)
		return errors.New(validatorFormatError(err))
	}
	return nil
}

func validatorFormatError(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		fieldname := err.Field()

		switch err.Tag() {
		case "required":
			return "El campo " + fieldname + " es requerido"
		case "email":
			return "El campo " + fieldname + " no es un email válido"
		case "gte":
			return "El campo " + fieldname + " debe ser mayor o igual a " + err.Param()
		case "lte":
			return "El campo " + fieldname + " debe ser menor o igual a " + err.Param()
		case "gt":
			return "El campo " + fieldname + " debe ser mayor que " + err.Param()
		case "lt":
			return "El campo " + fieldname + " debe ser menor que " + err.Param()
		case "min":
			return "El campo " + fieldname + " debe ser mayor o igual a " + err.Param()
		case "max":
			return "El campo " + fieldname + " debe ser menor o igual a " + err.Param()
		case "eqfield":
			return "El campo " + fieldname + " debe ser igual al campo " + err.Param()
		case "nefield":
			return "El campo " + fieldname + " no debe ser igual al campo " + err.Param()
		case "alpha":
			return "El campo " + fieldname + " solo debe contener letras"
		case "alphanum":
			return "El campo " + fieldname + " solo debe contener letras y números"
		case "numeric":
			return "El campo " + fieldname + " solo debe contener números"
		case "hexadecimal":
			return "El campo " + fieldname + " solo debe contener caracteres hexadecimales"
		case "hexcolor":
			return "El campo " + fieldname + " solo debe contener un color hexadecimal"
		case "rgb":
			return "El campo " + fieldname + " solo debe contener un color RGB"
		case "rgba":
			return "El campo " + fieldname + " solo debe contener un color RGBA"
		case "hsl":
			return "El campo " + fieldname + " solo debe contener un color HSL"
		case "hsla":
			return "El campo " + fieldname + " solo debe contener un color HSLA"
		case "e164":
			return "El campo " + fieldname + " solo debe contener un número de teléfono E164"
		case "url":
			return "El campo " + fieldname + " solo debe contener una URL"
		case "uri":
			return "El campo " + fieldname + " solo debe contener un URI"
		case "base64":
			return "El campo " + fieldname + " solo debe contener un string base64"
		case "contains":
			return "El campo " + fieldname + " debe contener el valor " + err.Param()
		case "containsany":
			return "El campo " + fieldname + " debe contener al menos uno de los valores " + err.Param()
		case "excludes":
			return "El campo " + fieldname + " no debe contener el valor " + err.Param()
		case "excludesall":
			return "El campo " + fieldname + " no debe contener ninguno de los valores " + err.Param()
		case "excludesrune":
			return "El campo " + fieldname + " no debe contener el caracter " + err.Param()
		case "iscolor":
			return "El campo " + fieldname + " debe ser un color"
		case "oneof":
			return "El campo " + fieldname + " debe ser uno de los siguientes valores " + err.Param()
		case "isbn":
			return "El campo " + fieldname + " debe ser un ISBN válido"
		case "isbn10":
			return "El campo " + fieldname + " debe ser un ISBN10 válido"
		case "isbn13":
			return "El campo " + fieldname + " debe ser un ISBN13 válido"
		case "uuid":
			return "El campo " + fieldname + " debe ser un UUID válido"
		case "uuid3":
			return "El campo " + fieldname + " debe ser un UUID3 válido"
		case "uuid4":
			return "El campo " + fieldname + " debe ser un UUID4 válido"
		case "uuid5":
			return "El campo " + fieldname + " debe ser un UUID5 válido"
		case "ascii":
			return "El campo " + fieldname + " debe ser un ASCII válido"
		case "printascii":
			return "El campo " + fieldname + " debe ser un ASCII imprimible válido"
		case "multibyte":
			return "El campo " + fieldname + " debe ser un multibyte válido"
		case "datauri":
			return "El campo " + fieldname + " debe ser un data URI válido"
		case "latitude":
			return "El campo " + fieldname + " debe ser una latitud válida"
		case "longitude":
			return "El campo " + fieldname + " debe ser una longitud válida"
		case "ssn":
			return "El campo " + fieldname + " debe ser un SSN válido"
		case "ipv4":
			return "El campo " + fieldname + " debe ser una dirección IPv4 válida"
		case "ipv6":
			return "El campo " + fieldname + " debe ser una dirección IPv6 válida"
		case "ip":
			return "El campo " + fieldname + " debe ser una dirección IP válida"
		case "cidr":
			return "El campo " + fieldname + " debe ser un CIDR válido"
		case "cidrv4":
			return "El campo " + fieldname + " debe ser un CIDRv4 válido"
		case "cidrv6":
			return "El campo " + fieldname + " debe ser un CIDRv6 válido"
		case "tcp4_addr":
			return "El campo " + fieldname + " debe ser una dirección TCP4 válida"
		case "tcp6_addr":
			return "El campo " + fieldname + " debe ser una dirección TCP6 válida"
		case "tcp_addr":
			return "El campo " + fieldname + " debe ser una dirección TCP válida"
		case "udp4_addr":
			return "El campo " + fieldname + " debe ser una dirección UDP4 válida"
		case "udp6_addr":
			return "El campo " + fieldname + " debe ser una dirección UDP6 válida"
		case "udp_addr":
			return "El campo " + fieldname + " debe ser una dirección UDP válida"
		case "ip4_addr":
			return "El campo " + fieldname + " debe ser una dirección IP4 válida"
		case "ip6_addr":
			return "El campo " + fieldname + " debe ser una dirección IP6 válida"
		case "ip_addr":
			return "El campo " + fieldname + " debe ser una dirección IP válida"
		case "unix_addr":
			return "El campo " + fieldname + " debe ser una dirección UNIX válida"
		case "mac":
			return "El campo " + fieldname + " debe ser una dirección MAC válida"
		case "hostname":
			return "El campo " + fieldname + " debe ser un hostname válido"
		case "fqdn":
			return "El campo " + fieldname + " debe ser un FQDN válido"
		case "unique":
			return "El campo " + fieldname + " debe ser único"
		case "isdefault":
			return "El campo " + fieldname + " debe ser el valor por defecto"
		case "requiredif":
			return "El campo " + fieldname + " es requerido si el campo " + err.Param() + " tiene el valor " + err.Value().(string)
		case "requiredunless":
			return "El campo " + fieldname + " es requerido a menos que el campo " + err.Param() + " tenga el valor " + err.Value().(string)
		case "requiredwith":
			return "El campo " + fieldname + " es requerido si el campo " + err.Param() + " está presente"
		case "requiredwithall":
			return "El campo " + fieldname + " es requerido si el campo " + err.Param() + " está presente"
		case "requiredwithout":
			return "El campo " + fieldname + " es requerido si el campo " + err.Param() + " no está presente"
		case "requiredwithoutall":
			return "El campo " + fieldname + " es requerido si el campo " + err.Param() + " no está presente"
		case "excludedif":
			return "El campo " + fieldname + " no debe estar presente si el campo " + err.Param() + " tiene el valor " + err.Value().(string)
		case "excludedunless":
			return "El campo " + fieldname + " no debe estar presente a menos que el campo " + err.Param() + " tenga el valor " + err.Value().(string)
		case "excludedwith":
			return "El campo " + fieldname + " no debe estar presente si el campo " + err.Param() + " está presente"
		case "excludedwithall":
			return "El campo " + fieldname + " no debe estar presente si el campo " + err.Param() + " está presente"
		case "excludedwithout":
			return "El campo " + fieldname + " no debe estar presente si el campo " + err.Param() + " no está presente"
		case "excludedwithoutall":
			return "El campo " + fieldname + " no debe estar presente si el campo " + err.Param() + " no está presente"
		case "isfile":
			return "El campo " + fieldname + " debe ser un archivo"
		case "isdir":
			return "El campo " + fieldname + " debe ser un directorio"
		case "isjson":
			return "El campo " + fieldname + " debe ser un JSON válido"
		default:
			return "El campo " + fieldname + " no es válido"
		}
	}

	return ""
}

func BodyParser(ctx *fiber.Ctx, data interface{}) error {
	if err := ctx.BodyParser(data); err != nil {
		log.Println(err)
		return errors.New("El cuerpo de la petición no es válido")
	}

	if err := validateSchema(data); err != nil {
		return err
	}

	return nil
}
