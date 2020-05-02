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
	router.HandleFunc("/", controllers.HomeHandler)

	http.ListenAndServe(":8000", router)

}
