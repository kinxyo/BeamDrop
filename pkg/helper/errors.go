package helper

import (
	"fmt"
	"net/http"
)

func ErrorCatch(w http.ResponseWriter, e error, t string,  desc string) {
	if e != nil {

		fmt.Println("Error encountered | type: ", t)
		fmt.Println(desc)

		/* 
		if this was `rust`
		I would have created enums 
		Will refactor this someday 
		to use `iota` maybe. 
		*/

		switch t {

		case "internal":
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return

		case "nofile":
			http.Error(w, "File not found", http.StatusNotFound)
			return

		//add more case here ⬇️

		}
	}
}
