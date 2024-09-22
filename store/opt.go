package store

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/biangacila/biatechauth1/internal/utils"
	"sync"
	"time"
)

type Otp struct {
	Value     string    `json:"value"`
	Email     string    `json:"username"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type ServeOtp struct {
	sync.Mutex
	list map[string]Otp
}

var storeOtp *ServeOtp

func NewOtpStore() *ServeOtp {
	return &ServeOtp{
		list: make(map[string]Otp),
	}
}

func InitialOtpStore() {
	if storeOtp == nil {
		storeOtp = NewOtpStore()
	}
}
func GetStoreOtp() *ServeOtp {
	return storeOtp
}
func (s *ServeOtp) Get(otp string) (Otp, error) {
	s.Lock()
	defer s.Unlock()
	o, ok := s.list[otp]
	if !ok {
		return Otp{}, errors.New("otp not found")
	}
	if time.Now().After(o.ExpiredAt) {
		fmt.Println("###### ", o.ExpiredAt.String(), " > ", time.Now().String())
		delete(s.list, otp)
		return Otp{}, errors.New("otp expired")
	}
	return o, nil
}
func (s *ServeOtp) Remove(otp string) error {
	s.Lock()
	defer s.Unlock()
	delete(s.list, otp)
	return nil
}
func (s *ServeOtp) Set(email, value string) error {
	s.Lock()
	defer s.Unlock()
	otp := Otp{
		Value:     value,
		Email:     email,
		ExpiredAt: utils.GetExpiredAt(2),
	}
	s.list[otp.Value] = otp
	return nil
}
func (s *ServeOtp) InStore(otp string) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.list[otp]
	return ok
}
func (s *ServeOtp) Generate() string {
	otpLength := 6
	otpChars := "0123456789"
	randBytes := make([]byte, otpLength)
	_, err := rand.Read(randBytes)
	if err != nil {
		// Handle error
		panic(err)
	}
	var otp string
	for _, b := range randBytes {
		idx := int(b) % len(otpChars)
		otp += string(otpChars[idx])
	}
	return otp
}
