package services

import (
	"github.com/punkestu/hello-oauth/internal/pages"
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
