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

//GetRapers controller
//Function that handles the "/raper/get" endpoint
func GetRapers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Rap API][GET][RAPER][/raper/get]")

	rapers, err := models.GetRapers()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[Rap API][GET][RAPER][/raper/get][PASSED]")
	responses.JSON(w, http.StatusOK, rapers)
}

//GetRaperByID controller
//Function that handles the "/raper/get/{id}" endpoint
func GetRaperByID(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	fmt.Printf("[Rap API][GET][RAPER][/raper/get/{%s}]\n", idString)

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}

	raper, err := models.GetRaperByID(idInt)

	fmt.Printf("[Rap API][GET][RAPER][/raper/get/{%s}][PASSED]", idString)
	fmt.Println("[RAPERGET]:", raper)
	responses.JSON(w, http.StatusOK, raper)
}

//CreateRaper controller
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

	err = models.RaperValidation(raper)
	if err != nil {
		responseString := fmt.Sprintf("error: invalidad request format")
		responses.JSON(w, http.StatusBadRequest, responseString)
		return
	}

	rowsAffected, err := models.CreateRaper(raper)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Raper created:", raper)

	responseString := fmt.Sprintf("Affected rows: %d", rowsAffected)
	fmt.Print("[Rap API][POST][RAPER][/raper/post}][PASSED]")
	fmt.Println("[RAPERPOST]:", raper)
	responses.JSON(w, http.StatusCreated, responseString)
}

//UpdateRaper controller
//Function that allows to update a raper in the database
func UpdateRaper(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Rap API][PUT][RAPER][/raper/update]")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	raper := models.Raper{}

	err = json.Unmarshal(body, &raper)
	if err != nil {
		fmt.Println(err)
	}

	err = models.RaperValidation(raper)
	if err != nil {
		responseString := fmt.Sprintf("error: invalid request format")
		responses.JSON(w, http.StatusBadRequest, responseString)
		return
	}

	rowsAffected, err := models.UpdateRaper(raper)
	if err != nil {
		fmt.Println(err)
	}

	responseString := fmt.Sprintf("Affected rows: %d", rowsAffected)

	fmt.Print("[Rap API][PUT][RAPER][/raper/update][PASSED]")
	fmt.Println("[RAPERPUT]:", raper)
	responses.JSON(w, http.StatusCreated, responseString)
}

//DeleteRaper controller
//Function that allows to delete a register in teh database
func DeleteRaper(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	fmt.Printf("[Rap API][DELETE][RAPER][/raper/delete/{%s}]\n", idString)

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}

	raper, err := models.GetRaperByID(idInt)
	if err != nil {
		fmt.Println(err)
	}

	rowsAffected, err := models.DeleteRaper(idInt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Raper deleted:", raper)
	responseString := fmt.Sprintf("Affected rows: %d", rowsAffected)

	fmt.Printf("[Rap API][DELETE][RAPER][/raper/delete/{%s}][PASSED]", idString)
	fmt.Println("[RAPERDELETE]:", raper)
	responses.JSON(w, http.StatusCreated, responseString)
}
