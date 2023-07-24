package services

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"net/http"
)

// config (you get all of it from Google console)
var oauthConfFb = &oauth2.Config{
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURL:  "http://localhost:8080/authorizedfb", // a link to redirect when success
	Scopes:       []string{"public_profile", "email"},
	Endpoint:     facebook.Endpoint,
}

func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfFb)
}

func CallBackFromFacebook(w http.ResponseWriter, r *http.Request) {
	GetUserData(w, r, oauthConfFb, "https://graph.facebook.com/me")
}
