package handler

import (
	"html/template"
	"net/http"
    "github.com/kinxyo/BeamDrop/pkg/helper"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    
    tmpl, err := template.ParseFiles("web/templates/index.html")
    helper.ErrorCatch(w, err, "internal", "Coulnd't start the web due to parsing issues.")

    err = tmpl.Execute(w,nil)
    helper.ErrorCatch(w, err, "internal", "Coulnd't start the web due to failed template execution")
}
