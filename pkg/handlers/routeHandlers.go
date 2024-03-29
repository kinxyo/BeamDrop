package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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

func BeamDrop(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	if r.MultipartForm != nil {
		files := r.MultipartForm.File["beam"]
		for i := range files {
			file, err := files[i].Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			fmt.Printf("File Name: %s\n", files[i].Filename)
			fmt.Printf("File Size: %d\n", files[i].Size)
			fmt.Printf("File Header: %s\n", files[i].Header.Get("Content-Type"))

			if _, err := os.Stat("Storage"); os.IsNotExist(err) {
				err := os.MkdirAll("Storage", 0755)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			dst, err := os.Create("Storage/" + files[i].Filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}
