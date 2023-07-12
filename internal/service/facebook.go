package services

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io"
	"net/http"
	"net/url"
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

// CallBackFromFacebook take data from login url
func CallBackFromFacebook(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	println("CODE>>", code)

	if code == "" { // if code is empty then the user is not found and login is failed
		println("Code not found..")
		_, err := w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		if err != nil {
			println("ERROR>>", err.Error())
		}
		reason := r.FormValue("error_reason")
		if reason == "user_denied" { // only if a user is denied to log in
			_, err := w.Write([]byte("User has denied Permission.."))
			if err != nil {
				println("ERROR>>", err.Error())
			}
		}
	} else {
		token, err := oauthConfFb.Exchange(context.Background(), code) // get token from code
		if err != nil {
			println("oauthConfFb.Exchange() failed with "+"ERROR>>", err.Error()+"\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		println("TOKEN>> AccessToken>> " + token.AccessToken)
		println("TOKEN>> Expiration Time>> " + token.Expiry.String())
		if token.Valid() {
			println("TOKEN>> Valid>> YES")
		} else {
			println("TOKEN>> Valid>> NO")
		}

		resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
			url.QueryEscape(token.AccessToken)) // get user info
		if err != nil {
			println("Get: "+"ERROR>>", err.Error()+"\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		response, err := io.ReadAll(resp.Body)
		if err != nil {
			println("ReadAll: "+"ERROR>>", err.Error()+"\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		println("parseResponseBody: " + string(response) + "\n")

		_, err = w.Write([]byte(string(response)))
		if err != nil {
			println("ERROR>>", err.Error())
		}
		return
	}
}
