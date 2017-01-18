package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	authenticated := false
	if !authenticated {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	RenderTemplate(w, r, "index/home", nil)
}

func HandleUserNew(w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params) {
	RenderTemplate(w, r, "user/new", nil)
}

func HandleUserCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// retrive data from FormValue()
	user, err := NewUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		if IsValidationError(err) {
			RenderTemplate(w, r, "user/new", map[string]interface{}{
				"Error": err.Error(),
				"User":  user,
			})
			return
		}
		// panic if bcrypt went wrong
		log.Fatalln(err.Error)
		panic(err)
	}
	err = globalUserStore.Save(user)
	if err != nil {
		panic(err)
	}

	session := NewSession(w)
	session.UserID = user.ID
	// Debug
	// log.Printf("session created: %#v", session)
	err = globalSessionStore.Save(session)

	log.Println(session)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/?flash=User+Created", http.StatusFound)
}

func HandleImageNew(w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params) {
	//RenderTemplate
}
