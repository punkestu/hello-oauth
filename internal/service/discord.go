package services

import (
	"context"
	"golang.org/x/oauth2"
	"io"
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

// CallBackFromDiscord take data from login url
func CallBackFromDiscord(w http.ResponseWriter, r *http.Request) {
	println("discord")
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
		token, err := oauthConfDc.Exchange(context.Background(), code) // get token from code
		if err != nil {
			println("oauthConfDc.Exchange() failed with "+"ERROR>>", err.Error()+"\n")
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
		resp, err := oauthConfDc.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
		if err != nil {
			println("error")
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
