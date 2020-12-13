package model

import "github.com/deathwofl/wine-reviews/pkg/validator"

func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.IsEmail("email", r.Email)
	v.Required("email", r.Email)

	v.Required("password", r.Password)
	v.MinLenght("password", r.Password, 8)
	v.Required("confirmpassword", r.Confirmpassword)
	v.MinLenght("confirmpassword", r.Confirmpassword, 8)
	v.EqualToField("password", r.Password, "confirmpassword", r.Confirmpassword)

	v.Required("username", r.Username)
	v.MinLenght("username", r.Username, 6)

	v.Required("firstname", r.FirstName)
	v.MinLenght("firstname", r.FirstName, 2)

	v.Required("lastname", r.LastName)
	v.MinLenght("lastname", r.LastName, 2)

	return v.IsValid(), v.Errors
}

func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", l.Email)
	v.IsEmail("email", l.Email)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}
