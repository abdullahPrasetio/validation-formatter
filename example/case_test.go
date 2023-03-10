/********************************************************************************
* Temancode Example Test Package                                            	*
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package example

import (
	"fmt"
	"testing"

	"github.com/abdullahPrasetio/validation-formatter/utils/strcase"
)

func TestToSnackCase(t *testing.T) {
	halo := "Halo_Ya"
	fmt.Println(strcase.SnakeCaseToCamelCase(halo))
}

func TestAddSpaceCamelCase(t *testing.T) {
	halo := "HaloYa"
	fmt.Println(strcase.AddSpaceCamelCase(halo))
}

func TestAddSpaceSnackCase(t *testing.T) {
	halo := "Halo_Ya_ya_te"
	fmt.Println(strcase.AddSpaceSnackCase(halo))
}

// func TestToSnackCase(t *testing.T) {
// 	validateStruct()
// }
