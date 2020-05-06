package api

import (
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiRap/api/controllers"
	"github.com/DixonOrtiz/ApiRap/api/middlewares"
	"github.com/gorilla/mux"
)

//Run function
//Function that raises the server and define the endpoints
func Run() {
	router := mux.NewRouter()

	//Auth (Google OAuth2 and JWT) routes
	router.HandleFunc("/login", controllers.HandleGoogleLogin)
	router.HandleFunc("/callback", controllers.HandleGoogleCallback)

	//Raper routes
	router.HandleFunc("/raper/get", middlewares.SetMiddlewareAuthentication(controllers.GetRapers)).Methods("GET")              //debería pedir autenticación lector
	router.HandleFunc("/raper/get/{id}", middlewares.SetMiddlewareAuthentication(controllers.GetRaperByID)).Methods("GET")      //debería pedir autenticación lector
	router.HandleFunc("/raper/create", middlewares.SetMiddlewareAuthentication(controllers.CreateRaper)).Methods("POST")        //debería pedir autenticación de administrador
	router.HandleFunc("/raper/update", middlewares.SetMiddlewareAuthentication(controllers.UpdateRaper)).Methods("PUT")         //debería pedir autenticación de administrador
	router.HandleFunc("/raper/delete/{id}", middlewares.SetMiddlewareAuthentication(controllers.DeleteRaper)).Methods("DELETE") //debería pedir autenticación de administrador

	fmt.Println("Corriendo en el puerto 8000")
	http.ListenAndServe(":8000", router)

}
