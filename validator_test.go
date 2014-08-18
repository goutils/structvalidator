package structvalidator

import (
	"testing"
	"fmt"
)

type ZeroStruct struct {
	A int `required:"true"`
}

type StructWithRegex struct {
	A string `required:"true" match:"^[a-b0-9]*$"`
}

type StructWithOptionalFields struct {
	A int
	B string `required:true`
}

func TestValidateFailsForZeroFields(t *testing.T) {

	valid, errors := Validate(ZeroStruct{})
	if errors != nil && len(errors) != 0 && !valid {
		fmt.Println("Validator validate zero field")
		t.Fail()
	}
}

func TestValidateMatchesRegex(t *testing.T) {

	valid, errors := Validate(StructWithRegex{"abcd"})

	if errors != nil && !valid {
		fmt.Println("Matches given pattern")
		t.Fail()
	}
}

func TestValidateIgnoredOptionalFields(t *testing.T) {

	valid, errors := Validate(StructWithOptionalFields{B:"abcd"})

	if errors != nil && !valid {
		fmt.Println("Doesn't ignore optional fields")
		t.Fail()
	}
}
