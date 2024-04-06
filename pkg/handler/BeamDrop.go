package handler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/kinxyo/BeamDrop/pkg/helper"
)

func BeamDrop(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File requested from the server...")

    link := extractLink(r.URL)

    if beams, ok := beamMap[link]; ok {

		fmt.Println("-----------------")
		fmt.Println(beams)
		fmt.Println("-----------------")

        // Create a buffer to write our archive to
        buf := new(bytes.Buffer)

        // Create a new zip archive
        zipWriter := zip.NewWriter(buf)

        // Loop over the files
        for _, beam := range beams.File {

			file, err := os.Open("Storage/" + beam)
			err_desc := "Could not process for download: " + beam
			helper.ErrorCatch(w, err, "nofile", err_desc) //although this will never happen

            f, err := zipWriter.Create(beam)
            helper.ErrorCatch(w, err, "internal", "Could not zip the files")

			fmt.Printf("Copying %s to %s", file.Name(), f)
            io.Copy(f, file)

            file.Close()
        }

        // Close the archive
        zipWriter.Close()

        // Set the headers
        w.Header().Set("Content-Disposition", "attachment; filename="+link+".zip")
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Transfer-Encoding", "binary")

        // Write the zip to the response
        io.Copy(w, buf)
    } else {
        
		fmt.Println("Invalid link, redirecting to homepage...")

        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func extractLink(r *url.URL) string {

	return strings.TrimPrefix(r.Path, "/drop/")

}
