package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//HomeHandler function
//Function that handles the root endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	url := "http://country.io/capital.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}
