package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string
	Email          string
	HashedPassword string
	ID             string
}

const (
	passwordLength = 8
	hashCost       = 10
	userIDLength   //why declare this?
)

const (
	errNoUsernameEnum       = 1 << iota
	errNoEmailEnum          = 1 << iota
	errNoPasswordEnum       = 1 << iota
	errPasswardTooShortEnum = 1 << iota
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email:    email,
		Username: username,
	}
	errorcode := 0
	if username == "" {
		errorcode += errNoUsernameEnum
	}
	if email == "" {
		errorcode += errNoEmailEnum
	}
	if password == "" {
		errorcode += errNoPasswordEnum
	}
	if len(password) < passwordLength {
		errorcode += errPasswardTooShortEnum
	}

	if errorcode != 0 {
		reterr := ValidationError(errors.New(""))
		switch {
		case errorcode-errPasswardTooShortEnum > 0:
			reterr = ValidationError(errors.New(fmt.Sprintf("%s\n%s", errPasswardTooShort, reterr.Error())))
			errorcode -= errPasswardTooShortEnum
			fallthrough
		case errorcode-errNoPasswordEnum > 0:
			reterr = ValidationError(errors.New(fmt.Sprintf("%s\n%s", errNoPassword, reterr.Error())))
			errorcode -= errNoPasswordEnum
			fallthrough
		case errorcode-errNoEmailEnum > 0:
			reterr = ValidationError(errors.New(fmt.Sprintf("%s\n%s", errNoEmail, reterr.Error())))
			errorcode -= errNoEmailEnum
			fallthrough
		case errorcode-errNoUsernameEnum > 0:
			reterr = ValidationError(errors.New(fmt.Sprintf("%s\n%s", errNoUsername, reterr.Error())))
			errorcode -= errNoUsernameEnum
		}
		return user, reterr
	}

	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return user, err
	} else if existingUser != nil {
		return user, errUsernameExists
	}

	existingUser, err = globalUserStore.FindByEmail(email)
	if err != nil {
		return user, err
	} else if existingUser != nil {
		return user, errEmailExists
	}

	// If input is fine, hashed password and give it a ID.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	user.HashedPassword = string(hashedPassword)

	user.ID = GenerateID("user", userIDLength)

	return user, err //final err indicates bcrypt.GenerateFromPassword() went wrong
}
