package main 

import (
    "net/http"
    "log"
)

func main(){
    mux := http.NewServeMux()
    mux.Handle("/assets/",http.StripPrefix("/assets/",http.FileServer(http.Dir("assets/"))))
    mux.HandleFunc("/",homeRender)
    log.Fatal(http.ListenAndServe(":3000",mux))
}

func homeRender(w http.ResponseWriter, req *http.Request){
    RenderTemplate(w, req, "index/home",nil)
}