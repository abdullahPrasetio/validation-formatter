/********************************************************************************
* Temancode Lang Package                                            			*
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package lang

var LangID = map[string]string{
	"alpha":                ":attribute hanya boleh berisi huruf.",
	"alphanum":             ":attribute hanya boleh berisi huruf dan angka.",
	"alphanumunicode":      ":attribute hanya boleh berisi huruf, angka, strip, dan garis bawah.",
	"alphaunicode":         ":attribute hanya boleh berisi huruf dan unicode",
	"boolean":              ":attribute harus bernilai true atau false",
	"date":                 ":attribute bukan tanggal yang valid. format yang valid adalah :values",
	"email":                ":attribute harus berupa alamat surel yang valid.",
	"file":                 ":attribute harus berupa sebuah berkas.",
	"gt":                   ":attribute harus bernilai lebih besar dari :value.",
	"gte":                  ":attribute harus bernilai lebih besar dari atau sama dengan :value.",
	"ip":                   ":attribute harus berupa alamat IP yang valid.",
	"ipv4":                 ":attribute harus berupa alamat IPv4 yang valid.",
	"ipv6":                 ":attribute harus berupa alamat IPv6 yang valid.",
	"json":                 ":attribute harus berupa JSON string yang valid.",
	"lt":                   ":attribute harus bernilai kurang dari :value.",
	"lte":                  ":attribute harus bernilai kurang dari atau sama dengan :value.",
	"len":                  "The :attribute must not be greater than :param characters.",
	"mac":                  ":attribute harus berupa MAC address yang valid.",
	"max":                  ":attribute maksimal bernilai :param.",
	"min":                  ":attribute minimal berisi :values karakter.",
	"numeric":              ":attribute harus berupa angka.",
	"required":             ":attribute wajib diisi.",
	"required_if":          ":attribute wajib diisi bila :other adalah :value.",
	"required_unless":      ":attribute wajib diisi kecuali :other memiliki nilai :values.",
	"required_with":        ":attribute wajib diisi bila terdapat :values.",
	"required_with_all":    ":attribute wajib diisi bila terdapat :values.",
	"required_without":     ":attribute wajib diisi bila :values Kosong",
	"required_without_all": ":attribute wajib diisi bila :values Kosong Semua.",
	"unique":               ":attribute sudah ada sebelumnya.",
	"url":                  "Format :attribute tidak valid.",
	"uuid":                 ":attribute harus merupakan UUID yang valid.",
}
