package main

import "errors"

type ValidationError error

var (
    errNoUsername = ValidationError(errors.New("You must supply a username"))
    errNoEmail = ValidationError(errors.New("Your must supply an E-mail address"))
    errNoPassword = ValidationError(errors.New("You must supply a password"))
    errPasswardTooShort = ValidationError(errors.New("Your password is too short"))
)


// IsValidationError(error) (bool) uses "comma,ok" statement 
// to return if error is an ValidationError
func IsValidationError(err error) bool {
    _, ok := err.(ValidationError)
    return ok
}