package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiRap/api/auth"
	"github.com/DixonOrtiz/ApiRap/api/responses"
)

//SetMiddlewareAuthentication middleware
//This function verify that the jwt introduced is valid
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Rap API][MIDDLEWARE][RAPER][SetMiddlewareAuthentication]")

		err := auth.TokenValidRequest(r)

		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		fmt.Println("[Rap API][MIDDLEWARE][RAPER][SetMiddlewareAuthentication][PASSED]")
		next(w, r)
	}
}
