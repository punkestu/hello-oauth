package main

import (
	services "github.com/punkestu/hello-oauth/internal/service"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
