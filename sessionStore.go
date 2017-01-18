package main

import (
	"fmt"
)

type SessionStore interface {
	Find(string) (*Session, error)
	Save(*Session) error
	Delete(*Session) error
}

var globalSessionStore SessionStore

// first trail, file storage
func init() {
	ss, err := NewFileSessionStore("./data/sessions.json")
	if err != nil {
		panic(fmt.Errorf("Error creating session storage: %s", err))
	}
	globalSessionStore = ss
}
