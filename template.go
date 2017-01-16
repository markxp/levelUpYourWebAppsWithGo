package main

import (
    "net/http"
    "html/template"
    "fmt"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

func RenderTemplate(w http.ResponseWriter, req *http.Request, name string, data interface{}){
    err := templates.ExecuteTemplate(w, name, data)
    if err != nil {
        http.Error(w,fmt.Sprintf(errorTemplate,name,err),http.StatusInternalServerError)
    }
}

var errorTemplate =`
<html>
    <body>
        <h1>Error rendering template %s</h1>
        <p>%s</p>
    </body>
</html>
`