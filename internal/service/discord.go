package services

import (
	"golang.org/x/oauth2"
	"net/http"
)

// config (you get all of it from Google console)
var oauthConfDc = &oauth2.Config{
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURL:  "http://localhost:8080/authorizeddc", // a link to redirect when success
	Scopes:       []string{"identify"},
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://discord.com/api/oauth2/authorize",
		TokenURL:  "https://discord.com/api/oauth2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	},
}

func HandleDiscordLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfDc)
}

func CallBackFromDiscord(w http.ResponseWriter, r *http.Request) {
	GetUserData(w, r, oauthConfDc, "https://discord.com/api/users/@me")
}
