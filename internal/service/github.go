package services

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
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

// CallBackFromGithub take data from login url
func CallBackFromGithub(w http.ResponseWriter, r *http.Request) {
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
		token, err := oauthConfGh.Exchange(context.Background(), code) // get token from code
		if err != nil {
			println("oauthConfGh.Exchange() failed with "+"ERROR>>", err.Error()+"\n")
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

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		resp, err := client.Do(req)
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
