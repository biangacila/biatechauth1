package valueobjects

import (
	"fmt"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
	GivenName  string
	FamilyName string
	Phone      string
}

func GenerateAllTokens(email, givenName, familyName, phone string) (signedToken, signedFreshToken string, err error) {
	secretKey, err := utils.GetJwtSecretKey()
	if err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return "", "", err
	}

	expiresAt := time.Now().Add(time.Hour * time.Duration(48)).Unix()
	claims := SignedDetails{
		Email:      email,
		GivenName:  givenName,
		FamilyName: familyName,
		Phone:      phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	signedFreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))

	if err != nil {
		log.Println(err)
		return
	}
	return signedToken, signedFreshToken, err
}

// Function to get expiration time and validate a JWT token
func GetTokenExpiryAndValidity(tokenString string) (time.Time, bool, error) {
	secretKey, err := utils.GetJwtSecretKey()
	if err != nil {
		utils.NewLoggerSlog().Error(err.Error())
		return time.Time{}, false, err
	}
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the signing key
		return []byte(secretKey), nil
	})

	if err != nil {
		return time.Time{}, false, fmt.Errorf("error parsing token: %v", err)
	}

	// Check if the token is valid and extract the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get expiration time (exp claim)
		if exp, ok := claims["exp"].(float64); ok {
			// Convert to time.Time
			expirationTime := time.Unix(int64(exp), 0)

			// Check if the token has expired
			if time.Now().After(expirationTime) {
				return expirationTime, false, nil // Token is expired
			}
			return expirationTime, true, nil // Token is valid
		} else {
			return time.Time{}, false, fmt.Errorf("token does not have an expiration time")
		}
	}

	return time.Time{}, false, fmt.Errorf("invalid token")
}
