package auth

import (
  "time"
  jwt "github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret = []byte("postman_twitter_challenge")

// CreateJWTAuth generates a new JWT Auth Token.
func CreateJWTAuth() (string, error) {
	claims := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "yashish.postman",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(hmacSampleSecret)
}
