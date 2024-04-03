package handler

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    
	success := r.URL.Query().Get("success") == "true"

    tmpl, err := template.ParseFiles("web/templates/index.html")
    
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, struct {
        UploadSuccess bool
    }{
        UploadSuccess: success,
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
