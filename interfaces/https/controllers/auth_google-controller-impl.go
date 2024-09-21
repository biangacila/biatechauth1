package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/biangacila/biatechauth1/application/dtos"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/luvungula-go/global"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

var clientId, clientSecret, _ = utils.GetGoogleClientLoginWith()
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  constants.GOOGLE_CALLBACK_URL2,
	ClientID:     clientId,
	ClientSecret: clientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}
var sessionStore = make(map[string]*sessionData)

type sessionData struct {
	Token    *oauth2.Token
	UserInfo *dtos.UserDto
}

type AuthGoogleControllerWith struct {
	autService *services.AuthServiceImpl
}

func NewAuthController() *AuthGoogleControllerWith {
	return &AuthGoogleControllerWith{}
}
func (c *AuthGoogleControllerWith) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (c *AuthGoogleControllerWith) Login(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	hostRedirectUri := r.URL.Query().Get("redirect_uri")
	sessionId := r.URL.Query().Get("session_id") // todo get it from the client request
	if sessionId == "" {
		sessionId = uuid.New().String()
	}
	// Store the session ID in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionId,
		Path:  hostRedirectUri,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "session_uri",
		Value: hostRedirectUri,
		Path:  hostRedirectUri,
	})
	redirectUri := constants.GOOGLE_CALLBACK_URL2
	if strings.Contains(host, "localhost") {
		redirectUri = fmt.Sprintf("http://%v/backend-biatechdesk/api/auth/google/callback", host)
	} else if utils.ContainsIPAddress(host) {
		redirectUri = fmt.Sprintf("http://localhost:8080/backend-biatechdesk/api/auth/google/callback")
	}
	googleOauthConfig.RedirectURL = redirectUri
	url := googleOauthConfig.AuthCodeURL("random_state")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func (c *AuthGoogleControllerWith) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Fetch user info (optional)
	client := googleOauthConfig.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	userInfo := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&userInfo)
	//todo just to test again
	uInfo := dtos.UserDto{}
	json.NewDecoder(resp.Body).Decode(&uInfo)
	global.DisplayObject("1. ):( userInfo resp", userInfo)
	global.DisplayObject("2. ):( userInfo resp", uInfo)

	if sessionStore == nil {
		sessionStore = make(map[string]*sessionData)
	}

	newSessionId, _ := generateSessionID()
	err = c.autService.RegisterToken(newSessionId, token, "google", userInfo)
	if err != nil {
		fmt.Println("!> error > ", err)
		return
	}

	// Store with our location token register
	userCode := c.autService.StoreTokenRealtime(token.AccessToken, userInfo)
	// Retrieve the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session ID not found", http.StatusUnauthorized)
		return
	}
	sessionID := cookie.Value

	cookieUri, _ := r.Cookie("session_uri")
	sessionUri := cookieUri.Value

	uriRed := fmt.Sprintf("%v?token=%v&session_id=%v&user_code=%v", sessionUri, token.AccessToken, sessionID, userCode)
	http.Redirect(w, r, uriRed, http.StatusFound)
	return
}

func (c *AuthGoogleControllerWith) CallbackGothic(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, "Failed to complete user authentication", http.StatusInternalServerError)
		return
	}

	// Use the user info and token from Goth
	fmt.Println("User: ", user)

	// Store user info in session if needed
	err = c.autService.RegisterToken(user.UserID, &oauth2.Token{
		AccessToken: user.AccessToken,
	}, "google", map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	})
	if err != nil {
		fmt.Println("!> error > ", err)
		return
	}

	// Store with our location token register
	userCode := c.autService.StoreTokenRealtime(user.AccessToken, utils.ObjectToMap(user))

	// Redirect or process the user data
	cookie, err := r.Cookie("session_id")
	sessionID := cookie.Value
	cookieUri, _ := r.Cookie("session_uri")
	sessionUri := cookieUri.Value
	uriRed := fmt.Sprintf("%v?token=%v&session_id=%v&user_code=%v", sessionUri, user.AccessToken, sessionID, userCode)

	http.Redirect(w, r, uriRed, http.StatusFound)
}
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
