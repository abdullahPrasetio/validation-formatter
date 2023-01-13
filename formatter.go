/********************************************************************************
* Temancode Formatter Validation Package                                        *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package formatter

import (
	"errors"
	"strings"

	"github.com/abdullahPrasetio/validation-formatter/lang"
	"github.com/abdullahPrasetio/validation-formatter/utils/strcase"
	"github.com/go-playground/validator/v10"
)

type addspace func(string) string

type validateFormatter struct {
	Language lang.Language
	TypeKey  TypeCase
}

type ValidateFormatter interface {
	GetErrorMsgValidation(err error) map[string]any
}

func NewValidateFormatter(language lang.Language, typeCase TypeCase) *validateFormatter {
	newValidator(typeCase)
	return &validateFormatter{Language: language, TypeKey: typeCase}
}

func (vf *validateFormatter) GetErrorMsgValidation(err error) map[string]any {

	errorMessages := map[string]any{}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		//out := make([]ErrorMsg, len(ve))
		for _, fe := range ve {
			var errorMsg = []string{}
			if vf.TypeKey == "SnackCase" {
				errorMsg = append(errorMsg, messageReplacer(getErrorMsg(fe, vf.Language), fe, strcase.AddSpaceSnackCase))
				errorMessages[strcase.ToSnakeCase(fe.Field())] = errorMsg
			} else if vf.TypeKey == "CamelCase" {
				errorMsg = append(errorMsg, messageReplacer(getErrorMsg(fe, vf.Language), fe, strcase.AddSpaceCamelCase))
				errorMessages[strcase.SnakeCaseToCamelCase(fe.Field())] = errorMsg
			} else {
				errorMsg = append(errorMsg, messageReplacer(getErrorMsg(fe, vf.Language), fe, strcase.AddSpaceCamelCase))
				errorMessages[fe.Field()] = errorMsg
			}
		}
	}

	return errorMessages
}

func getErrorMsg(fe validator.FieldError, language lang.Language) string {

	messages := language.GetData()
	message := messages[fe.Tag()]
	if len(message) <= 0 {
		return "Unknown error"
	}

	return message
}

func messageReplacer(message string, fe validator.FieldError, addSpace addspace) (newMessage string) {
	newMessage = strings.Replace(message, ":attribute", addSpace(fe.Field()), 1)
	newMessage = strings.Replace(newMessage, ":values", addSpace(fe.Param()), 1)
	newMessage = strings.Replace(newMessage, ":param", addSpace(fe.Param()), 1)
	getParamValue := strings.Split(fe.Param(), " ")
	var value string
	var param string
	if len(getParamValue) > 1 {
		value = getParamValue[1]
		param = getParamValue[0]
	}
	newMessage = strings.Replace(newMessage, ":other", addSpace(param), 1)
	newMessage = strings.Replace(newMessage, ":value", addSpace(value), 1)

	return
}
