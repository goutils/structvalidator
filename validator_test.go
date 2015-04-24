package structvalidator

import "testing"

func TestValidateFailsForZeroFields(t *testing.T) {
	type ZeroStruct struct {
		A int `required:"true"`
	}

	valid, errors := Validate(ZeroStruct{})
	if errors != nil && len(errors) != 0 && valid {
		t.Log(errors)
		t.Fail()
	}
}

func TestValidateMatchesRegex(t *testing.T) {
	type StructWithRegex struct {
		A string `required:"true" match:"^[a-z0-9]*$"`
	}

	valid, errors := Validate(StructWithRegex{"a2bcd"})
	if !valid {
		t.Log(errors)
		t.Fail()
	}
}

func TestValidateOnlyRequiredFields(t *testing.T) {
	type Struct struct {
		A int
		B string `required:"true"`
	}

	valid, errors := Validate(Struct{B: "abcd"})
	if !valid {
		t.Log(errors)
		t.Fail()
	}
}

func TestValidatesSliceFields(t *testing.T) {
	type StructWithSliceFields struct {
		A []int `required:"true"`
	}

	valid, errors := Validate(StructWithSliceFields{})
	if valid {
		t.Log(errors)
		t.Fail()
	}
}

func TestValidateShouldIgnoreSpecifiedFields(t *testing.T) {
	type StructWithMandatoryFields struct {
		A int      `required:"true"`
		B string   `required:"true"`
		C string   `required:"true"`
		D []string `required:"true"`
	}

	valid, errs := Validate(StructWithMandatoryFields{A: 123}, "B", "C", "D")
	if !valid {
		t.Log(errs)
		t.Fail()
	}
}
