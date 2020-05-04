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

	router.HandleFunc("/raper/get", controllers.GetRapers).Methods("GET")              //debería pedir autenticación lector
	router.HandleFunc("/raper/get/{id}", controllers.GetRaperByID).Methods("GET")      //debería pedir autenticación lector
	router.HandleFunc("/raper/create", controllers.CreateRaper).Methods("POST")        //debería pedir autenticación de administrador
	router.HandleFunc("/raper/update", controllers.UpdateRaper).Methods("PUT")         //debería pedir autenticación de administrador
	router.HandleFunc("/raper/delete/{id}", controllers.DeleteRaper).Methods("DELETE") //debería pedir autenticación de administrador

	http.ListenAndServe(":8000", router)

}
