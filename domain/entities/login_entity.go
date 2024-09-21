package entities

import "time"

type Login struct {
	Username         string
	SignedToken      string
	SignedFreshToken string
	Provider         string
	Updated          time.Time
	ExpiredAt        time.Time
}
