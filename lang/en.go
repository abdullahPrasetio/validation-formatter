/********************************************************************************
* Temancode Lang Package                                            			*
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package lang

var LangEN = map[string]string{
	"alpha":                "The :attribute must only contain letters.",
	"alphanum":             "The :attribute must only contain letters and numbers.",
	"alphanumunicode":      "The :attribute must only contain letters, numbers, dashes and underscores.",
	"alphaunicode":         "The :attribute must only contain letters and unicode.",
	"boolean":              "The :attribute field must be true or false.",
	"date":                 "The :attribute does not match the format :values.",
	"email":                "The :attribute must be a valid email address.",
	"file":                 "The :attribute must be a file.",
	"gt":                   "The :attribute must be greater than :value.",
	"gte":                  "The :attribute must be greater than or equal to :value.",
	"ip":                   "The :attribute must be a valid IP address.",
	"ipv4":                 "The :attribute must be a valid IPv4 address.",
	"ipv6":                 "The :attribute must be a valid IPv6 address.",
	"json":                 "The :attribute must be a valid JSON string.",
	"lt":                   "The :attribute must be less than :value.",
	"lte":                  "The :attribute must be less than or equal to :value.",
	"len":                  "The :attribute must not be greater than :param characters.",
	"mac":                  "The :attribute must be a valid MAC address.",
	"max":                  "The :attribute must not be greater than :param.",
	"min":                  "The :attribute must be at least :param.",
	"numeric":              "The :attribute must be a number.",
	"required":             "The :attribute field is required.",
	"required_if":          "The :attribute field is required when :other is :value.",
	"required_unless":      "The :attribute field is required unless :other is in :values.",
	"required_with":        "The :attribute field is required when :values is present.",
	"required_with_all":    "The :attribute field is required when :values are present.",
	"required_without":     "The :attribute field is required when :values is not present.",
	"required_without_all": "The :attribute field is required when none of :values are present.",
	"unique":               "The :attribute has already been taken.",
	"url":                  "The :attribute must be a valid URL.",
	"uuid":                 "The :attribute must be a valid UUID.",
}
