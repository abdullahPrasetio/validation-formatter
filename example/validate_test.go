package example

import (
	"fmt"
	"testing"

	formatter "github.com/abdullahPrasetio/validation-formatter"
	"github.com/abdullahPrasetio/validation-formatter/lang"
)

// Address houses a users address information
type Address struct {
	Street    string `validate:"required" json:"street"`
	City      string `validate:"required" json:"city"`
	Planet    string `validate:"required" json:"planet"`
	Phone     string `validate:"required,min=50" json:"phone"`
	TestName  string `validate:"required" json:"test_name"`
	UpdatedAt string `validate:"date=yyyy-mm-dd h:i:s" json:"updated_at"`
}

var validateFormatter formatter.ValidateFormatter

func init() {
	lang.AddLang("ED", ED)
	newLanguage := lang.NewLanguage("ED").ChangeMessage("required", "Halo :attribute").ChangeMessage("min", ":attribute min berisi :values ya")
	// newLanguage.ChangeMessage("date", ":attribute bukan sesuatu1 yang valid. format yang valid adalah :values")
	validateFormatter = formatter.NewValidateFormatter(newLanguage, "CamelCase")
	// lang.ChangeMessage("date", ":attribute bukan sesuatu yang valid. format yang valid adalah :values")
}

func ED() {
	lang.LangMessage = map[string]string{
		"date": ":attribute bukan tanggal yang valid. format yang valid wkwkwk adalah :values",
	}
}

func TestValidate(t *testing.T) {
	validateStruct()
}

func validateStruct() {
	validate := formatter.Validate

	address := &Address{
		Street:    "Eavesdown Docks",
		Planet:    "Persphone",
		UpdatedAt: "2022-10-21 10:12",
		// TestName:  "Test Name",
		Phone: "None",
		City:  "Bekasi",
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(address)
	if err != nil {
		fmt.Println(validateFormatter.GetErrorMsgValidation(err))
		// fmt.Println(err)
	}
}
