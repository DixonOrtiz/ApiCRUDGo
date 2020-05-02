package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiRap/api/models"
)

//ShowRaper function
//Function that handles the "/raper/create" endpoint
func ShowRaper(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Rap API][GET][RAPER][/raper/show]")

	rapers, err := models.GetRapers()
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(rapers)
}
