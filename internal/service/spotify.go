package services

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

// config (you get all of it from Google console)
var oauthConfSp = &oauth2.Config{
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURL:  "http://localhost:8080/authorizedsp", // a link to redirect when success
	Scopes:       []string{"user-read-private", "user-read-email"},
	Endpoint:     spotify.Endpoint,
}

func HandleSpotifyLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfSp)
}

func CallBackFromSpotify(w http.ResponseWriter, r *http.Request) {
	GetUserData(w, r, oauthConfSp, "https://api.spotify.com/v1/me")
}
