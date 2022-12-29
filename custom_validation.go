/********************************************************************************
* Temancode Formatter Validation Package                                        *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package formatter

import (
	"reflect"
	"strings"

	"github.com/abdullahPrasetio/validation-formatter/custom/date"
	"github.com/abdullahPrasetio/validation-formatter/utils/strcase"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

type TypeCase string

const (
	SnackCase TypeCase = "SnackCase"
	CamelCase TypeCase = "CamelCase"
)

func newValidator(typeCase TypeCase) {
	Validate = validator.New()
	Validate.RegisterValidation("date", date.DateValidation)
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		if typeCase == "CamelCase" {
			name = strcase.SnakeCaseToCamelCase(name)
		} else {
			name = strcase.ToSnakeCase(name)
		}

		return name
	})
}
