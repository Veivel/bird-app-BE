package authlib

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/providers/google"
)

var key = "asdasjdajdasjda"
var cookieStore = sessions.NewCookieStore([]byte(key))

func GetGoogleProvider(clientId string, clientSecret string) (googleProvider *google.Provider) {
	googleProvider = google.New(
		clientId,
		clientSecret,
		fmt.Sprintf("%s/auth/oauth2/callback", os.Getenv("FE_BASE_URL")),
		"email", "profile",
	)

	return
}

func GetCookieStore(callbackUrl string) (store *sessions.CookieStore) {
	var maxAge = 86400 * 30 // 30 days
	var isProd = false      // Set to true when serving over https

	store = cookieStore
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	return
}
