package handler

import (
	"html/template"
	"net/http"
    "github.com/kinxyo/BeamDrop/core/helper"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    
    tmpl, err := template.ParseFiles("template/index.html")
    helper.ErrorCatch(w, err, "internal", "Coulnd't parse the `html`.")

    err = tmpl.Execute(w,nil)
    helper.ErrorCatch(w, err, "internal", "Coulnd't start the web due to failed template execution")
}
