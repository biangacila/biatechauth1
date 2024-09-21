package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func GoogleAuthCallbackUri() string {
	if os.Getenv("GOOGLE_CALLBACK_URL") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return ""
		}
	}
	return os.Getenv("GOOGLE_CALLBACK_URL")
}
func GetAllowedMethods() []string {
	if os.Getenv("ALLOWED_METHODS") != "" {
		return strings.Split(os.Getenv("ALLOWED_METHODS"), ",")
	}
	return []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}
}

func GetAllowedOrigins() []string {
	if os.Getenv("ORIGINS") != "" {
		return strings.Split(os.Getenv("ORIGINS"), ",")
	}
	return []string{
		"http://localhost:*",
		"https://cloudcalls.easipath.com",
	}
}
func GetAllowedHeaders() []string {
	if os.Getenv("ALLOWED_HEADERS") != "" {
		return strings.Split(os.Getenv("ALLOWED_HEADERS"), ",")
	}
	return []string{
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
		"Authorization", "Origin",
		"X-Requested-With",
		"Accept",
		"Content-Type",
		"user-code",
		"org-code",
	}
}

func GetJwtSecretKey() (string, error) {
	if os.Getenv("JWT_SECRET_KEY") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return "", fmt.Errorf("error loading .env file: %v", err)
		}
	}
	return os.Getenv("JWT_SECRET_KEY"), nil
}
func GetGoogleClientLoginWith() (clientId, clientSecret string, err error) {
	if os.Getenv("GOOGLE_CLIENT_ID") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return "", "", fmt.Errorf("error loading .env file: %v", err)
		}
	}

	clientId = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	return clientId, clientSecret, nil
}
