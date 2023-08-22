package services

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

// config (you get all of it from Google console)
var oauthConfSp = &oauth2.Config{
	ClientID:     "7983c266d5d44649a247d8e35dc2e41b",
	ClientSecret: "1a2fc7c227d041328c315dbc7f55ec2b",
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
