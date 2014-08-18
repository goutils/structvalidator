package structvalidator

import (
	"reflect"
	"errors"
	"regexp"
)

const (
	REQUIRED = "required"
	MATCH = "match"
)


func Validate(v interface{}) (bool, []error) {
	var validationErrors []error
	vType := reflect.TypeOf(v)
	value := reflect.ValueOf(v)

	for index := 0; index < vType.NumField(); index++ {

		field := vType.Field(index)
		fieldValue := value.FieldByName(field.Name)


		switch fieldValue.Kind() {

		case reflect.Struct:
			_, err := Validate(fieldValue.Interface())
			if err!=nil && len(err) != 0 {
				validationErrors = append(validationErrors, err...)
			}
		//TODO validate slice elements
		case reflect.Slice:
			status := field.Tag.Get(REQUIRED)
			if status == "true" {
				if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
					err := errors.New("Required Field " + field.Name)
					validationErrors = append(validationErrors, err)
				}
			}
		default:
			err := validateField(field, fieldValue)
			if err != nil {
				validationErrors = append(validationErrors, err)
			}
		}
	}

	if len(validationErrors) == 0 {
		return true, validationErrors
	} else {
		return false, validationErrors
	}

}

func validateField(field reflect.StructField, fieldValue reflect.Value) error{
	status := field.Tag.Get(REQUIRED)
	if status == "true"{
		if fieldValue.Interface() == reflect.Zero(fieldValue.Type()).Interface() {
			err := errors.New("Required Field " + field.Name)
			return err
		}
		regex := field.Tag.Get(MATCH)
		if regex != "" {
			str, ok := fieldValue.Interface().(string)

			if !ok {
				err := errors.New("Not a string " + field.Name)
				return err
			}
			matched, err := regexp.MatchString(regex, str)
			if err != nil || !matched {
				err := errors.New("Doesn't match pattern " + field.Name)
				return err
			}
		}
	}
	return nil
}


