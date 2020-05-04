package controllers

import (
	"fmt"
	"net/http"

	"github.com/DixonOrtiz/ApiRap/api/auth"
)

//HandleGoogleLogin controler
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := auth.GoogleOauthConfig.AuthCodeURL(auth.OauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//HandleGoogleCallback controller
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := auth.GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", content)
}
