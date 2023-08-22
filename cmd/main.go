package main

import (
	"log"
	"net/http"

	services "github.com/punkestu/hello-oauth/internal/service"
)

func main() {

	// Routes for the application
	http.HandleFunc("/", services.HandleMain)
	http.HandleFunc("/login-gl", services.HandleGoogleLogin)
	http.HandleFunc("/authorized", services.CallBackFromGoogle)

	http.HandleFunc("/login-gh", services.HandleGithubLogin)
	http.HandleFunc("/authorizedgh", services.CallBackFromGithub)

	http.HandleFunc("/login-fb", services.HandleFacebookLogin)
	http.HandleFunc("/authorizedfb", services.CallBackFromFacebook)

	http.HandleFunc("/login-dc", services.HandleDiscordLogin)
	http.HandleFunc("/authorizeddc", services.CallBackFromDiscord)

	http.HandleFunc("/login-sp", services.HandleSpotifyLogin)
	http.HandleFunc("/authorizedsp", services.CallBackFromSpotify)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
