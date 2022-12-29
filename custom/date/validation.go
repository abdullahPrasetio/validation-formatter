/********************************************************************************
* Temancode Custom Validation Date Package                                      *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package date

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func DateValidation(fl validator.FieldLevel) bool {
	result := false
	var regexString string
	regexStrings := Mapping
	regexString = regexStrings[fl.Param()]
	if len(regexString) <= 0 {
		regexString = regexStrings["yyyy/mm/dd"]
	}

	re := regexp.MustCompile(regexString)
	m := re.MatchString(fl.Field().String())
	if m {
		result = true
	}
	return result
}
