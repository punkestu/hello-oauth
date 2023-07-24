package services

import (
	"context"
	"github.com/punkestu/hello-oauth/internal/pages"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

/*
HandleMain Function renders the index page when the application index route is called
*/
func HandleMain(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(pages.IndexPage))
	if err != nil {
		println(err.Error())
	}
}

/*
HandleLogin Function
*/
func HandleLogin(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		println("ERROR>>", err.Error())
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	URL.RawQuery = parameters.Encode() // add all config to login url
	mUrl := URL.String()
	http.Redirect(w, r, mUrl, http.StatusTemporaryRedirect) // go to login url
}

func GetUserData(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config, userEndPoint string) {
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
		token, err := oauthConf.Exchange(context.Background(), code) // get token from code
		if err != nil {
			println("oauthConf.Exchange() failed with "+"ERROR>>", err.Error()+"\n")
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

		resp, err := oauthConf.Client(context.Background(), token).Get(userEndPoint)
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
