package authproviders

import (
	"fmt"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	Key    = "randomString"
	MaxAge = 86400 * 30
	IsPro  = false
)

func NewGoogleAuth() {
	googleClientID := clientId                         // os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := clientSecret                 // os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectURL := constants.GOOGLE_CALLBACK_URL // os.Getenv("GOOGLE_CALLBACK_URL")

	fmt.Println(googleClientID, googleClientSecret, googleRedirectURL)

	store := sessions.NewCookieStore([]byte(Key))
	store.Options = &sessions.Options{}
	store.Options.MaxAge = MaxAge
	store.Options.HttpOnly = true
	store.Options.Secure = IsPro
	store.Options.Path = "/"

	gothic.Store = store
	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, googleRedirectURL),
	)

}
