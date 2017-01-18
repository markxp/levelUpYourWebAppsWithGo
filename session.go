package main

import (
	"log"
	"net/http"
	"time"
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
}

const (
	sessionLength     = 1 * time.Hour
	sessionCookieName = "Gopher_Session" //Do Not use whitespace
	sessionIDLength   = 20
)

func NewSession(w http.ResponseWriter) *Session {
	expiry := time.Now().Add(sessionLength)

	s := &Session{
		ID:     GenerateID("session", sessionIDLength),
		Expiry: expiry,
	}

	c := http.Cookie{
		Name:    sessionCookieName,
		Value:   s.ID,
		Expires: expiry,
	}

	http.SetCookie(w, &c)
	// Debug
	// log.Printf("set cookie: %#v", c)
	return s
}
