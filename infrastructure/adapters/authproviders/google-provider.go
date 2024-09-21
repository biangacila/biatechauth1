package authproviders

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/biatechauth1/store"
	"github.com/biangacila/luvungula-go/global"
	"github.com/pborman/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
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
	UserInfo *User
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	hostRedirectUri := r.URL.Query().Get("redirect_uri")
	sessionId := r.URL.Query().Get("session_id") // todo get it from the client request
	if sessionId == "" {
		sessionId = uuid.New()
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
func HandleGoogleSuccess(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("token")
	fmt.Println("HandleGoogleSuccess >>>> ", code, " ><")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	log.Printf("OAuth 2.0 token: %+v", token)
	w.Write([]byte("OAuth 2.0 login successful!"))

}
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
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
	uInfo := User{}
	json.NewDecoder(resp.Body).Decode(&uInfo)
	global.DisplayObject("1. ):( userInfo resp", userInfo)
	global.DisplayObject("2. ):( userInfo resp", uInfo)

	if sessionStore == nil {
		sessionStore = make(map[string]*sessionData)
	}

	newSessionId, _ := generateSessionID()
	err = storeTokenAndUserInfo(newSessionId, token, userInfo)
	if err != nil {
		fmt.Println("!> error > ", err)
		return
	}

	// Store with our location token register
	if err = store.GetStore().AddToken(code, token.AccessToken, token.Expiry); err != nil {
		utils.NewLoggerSlog().Error(err.Error())
	}

	// Retrieve the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session ID not found", http.StatusUnauthorized)
		return
	}
	sessionID := cookie.Value

	cookieUri, _ := r.Cookie("session_uri")
	sessionUri := cookieUri.Value

	uriRed := fmt.Sprintf("%v?token=%v&session_id=%v&user_code=%v", sessionUri, token.AccessToken, sessionID, "any")
	http.Redirect(w, r, uriRed, http.StatusFound)
	return
}
func HandleGoogleValidateToken(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.URL.Query().Get("token")
	userInfo, err := ValidateTokenGet(tokenStr)
	if err != nil {
		http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := map[string]interface{}{
		"message": "Protected data",
		"user":    userInfo,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func ValidateTokenGet(tokenStr string) (*User, error) {
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, tokenStr)
	if err != nil {
		return nil, errors.New("failed to exchange token")
	}
	if !token.Valid() {
		return nil, err
	}
	for _, info := range sessionStore {
		if info.Token == token {
			return info.UserInfo, nil
		}
	}

	return &User{}, nil
}
func storeTokenAndUserInfo(sessionID string, token *oauth2.Token, uInfo interface{}) error {
	str, _ := json.Marshal(uInfo)
	var userInfo *User

	_ = json.Unmarshal(str, &userInfo)
	sessionStore[sessionID] = &sessionData{
		Token:    token,
		UserInfo: userInfo,
	}
	return nil
}

/*func registerTokenWithLocalStore(token string, userIn interface{}) string {
	str, _ := json.Marshal(userIn)
	var userInfo User
	_ = json.Unmarshal(str, &userInfo)

	username := userInfo.Email
	var service LoginService
	isFind, record := service.IsUsernameExist(username)
	record.Picture = userInfo.Picture

	if !isFind {
		global.DisplayObject("RegisterTokenWithLocalStore isFind", userIn)
		return ""
	}
	//todo let send out our success result
	store.StoreAuth.Add(token, record)
	return record.Code
}*/

func generateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
