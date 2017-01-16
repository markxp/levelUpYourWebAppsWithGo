package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func AuthenticateRequest(w http.ResponseWriter, r *http.Request){
    authenticated := false
    if !authenticated {
        http.Redirect(w, r, "/register", http.StatusFound)
    }
}

func HandleHome(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    RenderTemplate(w,r ,"index/home",nil)
}

func HandleUserNew(w http.ResponseWriter, 
    r *http.Request, 
    _ httprouter.Params){
        RenderTemplate(w,r,"user/new",nil)
}

func HandleImageNew(w http.ResponseWriter,
    r *http.Request,
    params httprouter.Params){
        //RenderTemplate
 }
