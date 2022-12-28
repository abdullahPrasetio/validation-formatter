package lang

var LangEN = map[string]string{
	"required":             "The :attribute field is required.",
	"required_if":          "The :attribute field is required when :other is :value.",
	"required_unless":      "The :attribute field is required unless :other is in :values.",
	"required_with":        "The :attribute field is required when :values is present.",
	"required_with_all":    "The :attribute field is required when :values are present.",
	"required_without":     "The :attribute field is required when :values is not present.",
	"required_without_all": "The :attribute field is required when none of :values are present.",
	"numeric":              "The :attribute must be a number.",
	"min":                  "The :attribute must be at least :param characters.",
	"date":                 "The :attribute is not a valid date.",
}
