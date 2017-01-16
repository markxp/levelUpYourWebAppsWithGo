package main 

import (
    "net/http"
    "log"
    "github.com/julienschmidt/httprouter"
)

func main(){
    router := httprouter.New()
    router.Handle("GET","/",HandleHome)
    router.ServeFiles("/assets/*filepath",http.Dir("assets/"))

    middleware := Middleware{}
    middleware.Add(router)

    log.Fatal(http.ListenAndServe(":3000",middleware))
}

