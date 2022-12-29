/********************************************************************************
* Temancode Strcase Package                                            			*
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package strcase

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchCamelCaseCap = regexp.MustCompile("(\\B[A-Z])")
var matchSnackCaseCap = regexp.MustCompile("(.)([A-Z][a-z]+)")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func AddSpaceSnackCase(str string) string {
	newText := strings.ReplaceAll(str, "_", " ")
	return newText
}
