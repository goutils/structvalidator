package structvalidator

import (
	"testing"
	"fmt"
)

type ZeroStruct struct {
	A int `required:"true"`
}

type StructWithRegex struct {
	A string `required:"true" match:"^[a-z0-9]*$"`
}

type StructWithOptionalFields struct {
	A int
	B string `required:"true"`
}

type StructWithSliceFields struct {
	A []int `required:"true"`
}

func TestValidateFailsForZeroFields(t *testing.T) {

	valid, errors := Validate(ZeroStruct{})
	if errors != nil && len(errors) != 0 && valid {
		fmt.Println(errors)
		t.Fail()
	}
}

func TestValidateMatchesRegex(t *testing.T) {

	valid, errors := Validate(StructWithRegex{"abcd"})
	if !valid {
		fmt.Println(errors)
		t.Fail()
	}
}

func TestValidateIgnoredOptionalFields(t *testing.T) {

	valid, errors := Validate(StructWithOptionalFields{B:"abcd"})

	if !valid {
		fmt.Println(errors)
		t.Fail()
	}
}

func TestValidatesSliceFields(t *testing.T) {

	valid, errors := Validate(StructWithSliceFields{})

	if valid {
		fmt.Println(errors)
		t.Fail()
	}
}
