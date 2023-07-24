package services

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

// config (you get all of it from Google console)
var oauthConfGl = &oauth2.Config{
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURL:  "http://localhost:8080/authorized", // a link to redirect when success
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfGl)
}

// CallBackFromGoogle take data from login url
func CallBackFromGoogle(w http.ResponseWriter, r *http.Request) {
	GetUserData(w, r, oauthConfGl, "https://www.googleapis.com/oauth2/v2/userinfo")
}
