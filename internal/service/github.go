package services

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
)

// config (you get all of it from Google console)
var oauthConfGh = &oauth2.Config{
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURL:  "http://localhost:8080/authorizedgh", // a link to redirect when success
	Scopes:       []string{"user"},
	Endpoint:     github.Endpoint,
}

func HandleGithubLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfGh)
}

func CallBackFromGithub(w http.ResponseWriter, r *http.Request) {
	GetUserData(w, r, oauthConfGh, "https://api.github.com/user")
}
