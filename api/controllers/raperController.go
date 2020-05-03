package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/DixonOrtiz/ApiRap/api/models"
	"github.com/DixonOrtiz/ApiRap/api/responses"
	"github.com/gorilla/mux"
)

//GetRapers function
//Function that handles the "/raper/get" endpoint
func GetRapers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Rap API][GET][RAPER][/raper/get]")

	rapers, err := models.GetRapers()
	if err != nil {
		fmt.Println(err)
	}

	responses.JSON(w, http.StatusOK, rapers)
}

//GetRaperByID function
//Function that handles the "/raper/get/{id}" endpoint
func GetRaperByID(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	fmt.Printf("[Rap API][GET][RAPER][/raper/get/{%s}]\n", idString)

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}

	raper, err := models.GetRaperByID(idInt)

	responses.JSON(w, http.StatusOK, raper)
}

//CreateRaper function
//Function that handles the "/raper/create" endpoint
func CreateRaper(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Rap API][POST][RAPER][/raper/create]")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	raper := models.Raper{}

	err = json.Unmarshal(body, &raper)
	if err != nil {
		fmt.Println(err)
	}

	affectedRows, err := models.CreateRaper(raper)
	if err != nil {
		fmt.Println(nil)
	}

	fmt.Println("Raper created:", raper)

	responseString := fmt.Sprintf("Affected rows: %d", affectedRows)
	responses.JSON(w, http.StatusCreated, responseString)
}
