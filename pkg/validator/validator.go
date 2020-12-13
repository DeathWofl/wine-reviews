package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

type Validator struct {
	Errors map[string]string
}

type Validation interface {
	Validate() (bool, map[string]string)
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// New create a new validator
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) MinLenght(field, value string, high int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) < high {
		v.Errors[field] = fmt.Sprintf("%s must be at last %d character long", field, high)
		return false
	}

	return true
}

func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if !emailRegex.MatchString(email) {
		v.Errors[field] = "Is not valid email"
		return false
	}

	return true
}

func (v *Validator) EqualToField(field string, value interface{}, toequalfield string, toequalvalue interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if value != toequalvalue {
		v.Errors[field] = fmt.Sprintf("%s must equal %s", field, toequalfield)
		return false
	}

	return true
}

func (v *Validator) Required(field, value string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if v.IsEmpty(value) {
		v.Errors[field] = fmt.Sprintf("%s is required", field)
		return false
	}

	return true
}

func (v *Validator) IsEmpty(value string) bool {
	t := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		return t.Len() == 0
	}

	return false
}
