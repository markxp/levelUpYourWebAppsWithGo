package main

// UserStore interface grants storage flexibility.
type UserStore interface {
    Find(string) (*User, error)
    FindByEmail(string) (*User, error)
    FindByUsername(string) (*User, error)
    Save(User) error
}

var globalUserStore UserStore

// first version, use UserStoreFile
func init(){
    store, err := NewFileUserStore("./data/user.json")
    if err != nil {
        // add extra information and invoke panic()
        panic(fmt.Errorf("Error creating user store: %s", err))
    }
    globalUserStore = store
}