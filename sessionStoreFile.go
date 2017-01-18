package main

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type FileSessionStore struct {
	filename string
	Sessions map[string]Session
}

func NewFileSessionStore(filename string) (*FileSessionStore, error) {
	f := &FileSessionStore{
		filename: filename,
		Sessions: map[string]Session{
		//empty initialization
		},
	}

	contents, err := ioutil.ReadFile(f.filename)
	if err != nil {
		// ignore empty file error
		if os.IsNotExist(err) {
			return f, nil
		}
		return nil, err
	}
	err = json.Unmarshal(contents, f) // same question, why unmarshal to f, not f.Sessions
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FileSessionStore) Find(id string) (*Session, error) {
	s, ok := f.Sessions[id]
	if !ok {
		return nil, nil // error for storage, not for empty data
	}
	return &s, nil
}

func (f *FileSessionStore) Save(s *Session) error {
	contents, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f.filename, contents, 0660)
}

func (f *FileSessionStore) Delete(s *Session) error {
	delete(f.Sessions, s.ID)
	contents, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f.filename, contents, 0660)
}
