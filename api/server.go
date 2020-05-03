package api

import (
	"net/http"

	"github.com/DixonOrtiz/ApiRap/api/controllers"
	"github.com/gorilla/mux"
)

//Run function
//Function that raises the server and define the endpoints
func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/raper/get", controllers.GetRapers).Methods("GET")
	router.HandleFunc("/raper/get/{id}", controllers.GetRaperByID).Methods("GET")
	router.HandleFunc("/raper/create", controllers.CreateRaper).Methods("POST")
	router.HandleFunc("/raper/update", controllers.UpdateRaper).Methods("PUT")

	http.ListenAndServe(":8000", router)

}
