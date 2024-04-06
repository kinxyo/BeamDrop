package handler

import (
	"github.com/kinxyo/BeamDrop/pkg/helper"
	. "github.com/kinxyo/BeamDrop/pkg/models"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"crypto/rand"
    "encoding/hex"
    "encoding/json"
)

var beamMap map[string]Beam = make(map[string]Beam)

func BeamUpload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request received...")

	var beam Beam;
	var filelist []string;

	r.ParseMultipartForm(32 << 20)

	if r.MultipartForm == nil {
		return
	}
	
	files := r.MultipartForm.File["beam"]
	
	/* Download files to the server */
	
	for i := range files {

		fmt.Printf("Initiating download: %d\n", i+1)

		file := files[i]
		name := file.Filename

		data, err := file.Open()
		helper.ErrorCatch(w, err, "internal", "Could not process the files sent to the server")
	
		defer data.Close()

		print_details(file)

		validate_storage(w)

		target, err := os.Create("Storage/" + name)
		err_desc := "Error truncating file: " + name
		helper.ErrorCatch(w, err, "internal", err_desc)
		
		defer target.Close()
		
		_, err = io.Copy(target, data) //🔴
		err_desc = "Error copying file: " + name
		helper.ErrorCatch(w, err, "internal", err_desc)


		
		filelist = append(filelist, name)
		
	}
	
	/* Generate link then associate the data */
	link := generateRandomLink()
	
	beam.Link = link
	beam.File = filelist

	fmt.Println("Going to store ⬇️")
	fmt.Println("link:", link)
	fmt.Println("filelist:", filelist)

	store(beam)

	// w.Write([]byte(beam.Link))
	linkData := map[string]string{"link": beam.Link}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(linkData)
}

func print_details(file *multipart.FileHeader) {

	fmt.Printf("File Name: %s\n", file.Filename)
	fmt.Printf("File Size: %d\n", file.Size)
	fmt.Printf("File Header: %s\n", file.Header.Get("Content-Type"))

}

func validate_storage(w http.ResponseWriter) {
	
	if _, err := os.Stat("Storage"); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll("Storage", 0755)
			helper.ErrorCatch(w, err, "internal", "Couldn't create directory for storage")
		} else {
			// handle other errors
			helper.ErrorCatch(w, err, "internal", "Error checking storage directory")
		}
	}

	fmt.Println("Storage health: 100%")

}

func generateRandomLink() string {
    byteSlice := make([]byte, 16)
    _, err := rand.Read(byteSlice)
    if err != nil {
        panic(err)
    }
    return hex.EncodeToString(byteSlice)
}

func store(beam Beam) {
	beamMap[beam.Link] = beam
}