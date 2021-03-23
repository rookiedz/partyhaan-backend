package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewToken() {}
func CreateToken(userID uint64, fullname string, authorized bool, types string) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized":    authorized,
		"user_id":       userID,
		"user_fullname": fullname,
		"exp":           time.Now().Add(15 * time.Minute),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
func ValidateToken() {}
