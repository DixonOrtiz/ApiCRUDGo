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

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/raper/show", controllers.ShowRaper).Methods("GET")

	http.ListenAndServe(":8000", router)

}
