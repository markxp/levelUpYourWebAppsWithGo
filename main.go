package main 

import (
    "net/http"
    "log"
    "github.com/julienschmidt/httprouter"
)

func main(){

    unauthenticatedRouter := httprouter.New()
    unauthenticatedRouter.GET("/", HandleHome)

    authenticatedRouter := httprouter.New()
    authenticatedRouter.GET("/images/new", HandleImageNew)

    middleware := Middleware{}
    middleware.Add(unauthenticatedRouter).
        Add(http.HandlerFunc(AuthenticateRequest)).
        Add(authenticatedRouter)


    log.Fatal(http.ListenAndServe(":3000",middleware))
}

