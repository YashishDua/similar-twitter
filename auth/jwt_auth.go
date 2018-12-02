package auth

import (
  "fmt"
  "errors"
  "time"
  "github.com/google/uuid"
  jwt "github.com/dgrijalva/jwt-go"
  "postman-twitter/models"
)

var hmacSecret = []byte("postman_twitter_challenge")

type JWTAuthInfo struct {
	UserID   *uuid.UUID `json:"user_id"`
  Username string     `json:"username"`
  Password string     `json"password"`
  IssuedAt time.Time  `json:"iat,omitempty"`
  jwt.StandardClaims
}

// CreateJWTAuth generates a new JWT Auth Token.
func CreateJWTAuth(userAuth models.UserAuth) (string, error) {
  claims := JWTAuthInfo{
    UserID: userAuth.ID,
    Username: userAuth.Username,
    Password: userAuth.Password,
    IssuedAt: time.Now(),
    StandardClaims: jwt.StandardClaims{
			Issuer:    "yashish-postman-twitter",
		},
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(hmacSecret)
}

// DecodeJWTAuth returns decoded JWTAuthInfo
func DecodeJWTAuth(tokenString string) (*JWTAuthInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTAuthInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	jwtAuthInfo, ok := token.Claims.(*JWTAuthInfo)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid JWT Auth Token")
	}
	return jwtAuthInfo, nil
}
