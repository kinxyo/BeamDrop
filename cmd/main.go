package main

import (
	"github.com/kinxyo/Beamsend/pkg/handlers"
	"net/http"
)

func main() {
	// Static files
	http.Handle("/web/static/assets/", http.StripPrefix("/web/static/assets/", http.FileServer(http.Dir("./web/static/assets"))))
	http.Handle("/web/static/css/", http.StripPrefix("/web/static/css/", http.FileServer(http.Dir("./web/static/css"))))
	http.Handle("/web/static/javascript/", http.StripPrefix("/web/static/javascript/", http.FileServer(http.Dir("./web/static/javascript"))))

	// Handle routes
	http.HandleFunc("/", handlers.Welcome)
	http.HandleFunc("/dropthebeam", handlers.BeamDrop)

	// Start the server
	http.ListenAndServe(":3000", nil)
}