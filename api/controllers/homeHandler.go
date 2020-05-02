package controllers

import (
	"fmt"
	"net/http"
)

//HomeHandler function
//Function that handles the root endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome home!")
}
