package main

// UserStore interface grants storage flexibility.
type UserStore interface {
    Find(string) (*User, error)
    FindByEmail(string) (*User, error)
    FindByUsername(string) (*User, error)
    Save(User) error
}