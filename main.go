package main

import (
	"github.com/kinxyo/BeamDrop/core/handler"
	"net/http"
)

func main() {
	// Static files
	const STATIC_PATH = "/support/"
	const ASSETS_PATH = "/assets/"

	http.Handle(STATIC_PATH, http.StripPrefix(STATIC_PATH, http.FileServer(http.Dir("template"+STATIC_PATH))))
    http.Handle(ASSETS_PATH, http.StripPrefix(ASSETS_PATH, http.FileServer(http.Dir("template"+ASSETS_PATH))))

	// Handle routes
	http.HandleFunc("/", handler.Welcome)
	http.HandleFunc("/beam", handler.BeamUpload)

	// Wildcard route
	http.HandleFunc("/drop/", handler.BeamDrop)

	// Start the server
	http.ListenAndServe(":3000", nil)
}
