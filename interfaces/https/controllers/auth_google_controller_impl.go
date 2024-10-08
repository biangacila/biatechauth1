package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/biangacila/biatechauth1/application/dtos"
	"github.com/biangacila/biatechauth1/application/services"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/luvungula-go/global"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strings"
)

var clientId, clientSecret, _ = utils.GetGoogleClientLoginWith()
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  utils.GoogleAuthCallbackUri(),
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

type AuthGoogleControllerImpl struct {
	service services.LoginService
}

func (c *AuthGoogleControllerImpl) ValidateToken(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewAuthGoogleController(
	service services.LoginService,
) *AuthGoogleControllerImpl {
	return &AuthGoogleControllerImpl{
		service: service,
	}
}
func (c *AuthGoogleControllerImpl) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (c *AuthGoogleControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	//host := r.Host
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

	redirectUri := utils.GoogleAuthCallbackUri()
	/*if strings.Contains(host, "localhost") {
		// TODO please uncomment
		redirectUri = fmt.Sprintf("http://%v/backend-biatechauth1/api/logins-google/callback", host)
	} else if utils.ContainsIPAddress(host) {
		redirectUri = fmt.Sprintf("http://localhost:8080/backend-biatechauth1/api/logins-google/callback")
	}*/
	googleOauthConfig.RedirectURL = redirectUri
	url := googleOauthConfig.AuthCodeURL("random_state")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func (c *AuthGoogleControllerImpl) Callback(w http.ResponseWriter, r *http.Request) {
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

	if sessionStore == nil {
		sessionStore = make(map[string]*sessionData)
	}

	//newSessionId, _ := generateSessionID()
	err = c.service.RegisterGoogleToken(token.AccessToken, utils.MapToString(userInfo))
	if err != nil {
		fmt.Println("!> error > ", err)
		return
	}

	// Retrieve the session ID from the cookie
	cookieId, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session ID not found", http.StatusUnauthorized)
		return
	}
	sessionID := cookieId.Value
	//sessionUri := extractValueFromQueryString(sessionID, "redirect_uri")
	global.DisplayObject("cookieId", cookieId)

	cookieUri, err := r.Cookie("session_uri")
	if err != nil {
		http.Error(w, "Session URL not found", http.StatusUnauthorized)
		return
	}
	sessionUri := cookieUri.Value
	fmt.Println(":)))sessionId>>>>", sessionID, " >>>>  ", sessionUri)
	global.DisplayObject("cookieUri", cookieUri)

	uriRed := fmt.Sprintf("%v?token=%v&session_id=%v&user_info=%v", sessionUri, token.AccessToken, sessionID, utils.MapToString(userInfo))

	http.Redirect(w, r, uriRed, http.StatusFound)
	return
}

func generateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func extractValueFromQueryString(input, key string) string {
	var value string
	// Split the query string
	if strings.Contains(input, "?") {
		// Extract the query string portion after '?'
		queryString := strings.Split(input, "?")[1]

		// Parse the query string into values
		params, err := url.ParseQuery(queryString)
		if err != nil {
			fmt.Println("Error parsing query string:", err)
			return value
		}

		// Extract the 'redirect_uri' value
		redirectURI := params.Get("redirect_uri")
		if redirectURI != "" {
			fmt.Println("Redirect URI:", redirectURI)
			value = redirectURI
		} else {
			fmt.Println("Redirect URI not found")
		}
	}
	return value
}
