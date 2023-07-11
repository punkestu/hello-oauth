package main

import (
	services "github.com/punkestu/hello-oauth/internal/service"
	"log"
	"net/http"
)

func main() {

	// Initialize Oauth2 Services
	services.InitializeOAuthGoogle()

	// Routes for the application
	http.HandleFunc("/", services.HandleMain)
	http.HandleFunc("/login-gl", services.HandleGoogleLogin)
	http.HandleFunc("/authorized", services.CallBackFromGoogle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}