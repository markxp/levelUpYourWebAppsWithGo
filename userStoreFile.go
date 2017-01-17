package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

//First step - save to a file with JSON format
type FileUserStore struct {
	filename string
	Users    map[string]User
}

func (f FileUserStore) Save(user User) error {
	f.Users[user.ID] = user
	contents, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.filename, contents, 0660)
	if err != nil {
		return err
	}
	return nil
}

// In Find() function series, type FileUserStore,
// error is not used yet. errors will be used to indicate storage/loading problem.
// If user == nil, which means NOT FOUND. If error != nil, storage/loading breaks.
func (f FileUserStore) Find(id string) (*User, error) {
	user, ok := f.Users[id]
	if ok {
		return &user, nil
	}
	return nil, nil
}

func (f FileUserStore) FindByUsername(username string) (*User, error) {
	// why filter empty username string?
	if username == "" {
		return nil, nil
	}

	for _, v := range f.Users {
		if strings.ToLower(username) == strings.ToLower(v.Username) {
			return &v, nil
		}
	}
	return nil, nil
}

func (f FileUserStore) FindByEmail(email string) (*User, error) {
	// why filter empty email string?
	if email == "" {
		return nil, nil
	}

	for _, v := range f.Users {
		if strings.ToLower(email) == strings.ToLower(v.Email) {
			return &v, nil
		}
	}
	return nil, nil
}

func NewFileUserStore(filename string) (*FileUserStore, error) {
	f := &FileUserStore{
		Users:    map[string]User{},
		filename: filename,
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		// file not exist is ok.
		if os.IsNotExist(err) {
			// return a new *FileUserStore
			return f, nil
		}
		return nil, err
	}
	// load file contents and transform into FileUserStore
	err = json.Unmarshal(contents, f) // strange, why not unmarshal to f.Users
	if err != nil {
		return nil, err
	}
	return f, nil
}
