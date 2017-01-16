package main

import (
    "net/http"
    "html/template"
    "fmt"
    "bytes"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

func RenderTemplate(w http.ResponseWriter, req *http.Request, name string, data interface{}){
    // overwrites "yield" function for "index/home"
    // write to buf and then return
    funcs := template.FuncMap{
        "yield": func() (template.HTML, error){
            buf := bytes.NewBuffer(nil)
            err := templates.ExecuteTemplate(buf,name,data)
            // use template.HTML() to make content safe-html
            // Note: template HTML treats all content as trusted html.
            //      Should not be used for third-party content
            return template.HTML(buf.String()),err
        },
    }

    // we do not render(Execute) layout directly. We do layoutClone.Execute()
    // Copy that template to keep FuncMap seperate. Since layoutClone is a new copy and register "funcs"
    // as its Funcs. 
    // Otherwise, using layout directly might accidently raise "layoutFuncs" for "layout"
    layoutClone, _ := layout.Clone()
    layoutClone.Funcs(funcs)
    err := layoutClone.Execute(w,data)

    if err != nil {
        http.Error(w,
         fmt.Sprintf(errorTemplate, name, err),
         http.StatusInternalServerError,
        )
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

// This piece of code is not actually used. It represents the FuncMap for "layout" template
// While in RenderTemplate(), we overwrites FuncMap.
var layoutFuncs = template.FuncMap{
    // Throw an error without doing any thing.
    "yield": func() (string,error) {
        return "",fmt.Errorf("yield called inappropriately")
    },
}

var layout = template.Must(template.New("layout.html").
            Funcs(layoutFuncs).
            ParseFiles("templates/layout.html"))