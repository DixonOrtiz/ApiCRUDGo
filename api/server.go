package api

import (
	"net/http"

	"github.com/DixonOrtiz/ApiRap/api/controllers"
	"github.com/DixonOrtiz/ApiRap/api/middlewares"
	"github.com/gorilla/mux"
)

//Run function
//Function that raises the server and define the endpoints
func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.HandleGoogleLogin)
	router.HandleFunc("/callback", controllers.HandleGoogleCallback)

	router.HandleFunc("/raper/get", controllers.GetRapers).Methods("GET")                                                  //debería pedir autenticación lector
	router.HandleFunc("/raper/get/{id}", middlewares.SetMiddlewareAuthentication(controllers.GetRaperByID)).Methods("GET") //debería pedir autenticación lector
	router.HandleFunc("/raper/create", controllers.CreateRaper).Methods("POST")                                            //debería pedir autenticación de administrador
	router.HandleFunc("/raper/update", controllers.UpdateRaper).Methods("PUT")                                             //debería pedir autenticación de administrador
	router.HandleFunc("/raper/delete/{id}", controllers.DeleteRaper).Methods("DELETE")                                     //debería pedir autenticación de administrador

	http.ListenAndServe(":8000", router)

}
