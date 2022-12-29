/********************************************************************************
* Temancode Lang Package                                            			*
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package lang

var LangID = map[string]string{
	"required":             ":attribute wajib diisi.",
	"required_if":          ":attribute wajib diisi bila :other adalah :value.",
	"required_unless":      ":attribute wajib diisi kecuali :other memiliki nilai :values.",
	"required_with":        ":attribute wajib diisi bila terdapat :values.",
	"required_with_all":    ":attribute wajib diisi bila terdapat :values.",
	"required_without":     ":attribute wajib diisi bila :values Kosong",
	"required_without_all": ":attribute wajib diisi bila :values Kosong Semua.",
	"numeric":              ":attribute harus berupa angka.",
	"min":                  ":attribute minimal berisi :values karakter.",
	"date":                 ":attribute bukan tanggal yang valid. format yang valid adalah :values",
}
