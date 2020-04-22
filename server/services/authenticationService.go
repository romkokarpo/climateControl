package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationService struct {
	jwtKey string

	JwtToken       string
	ExpirationTime time.Time
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{
		jwtKey: "Some random key for the time being",
	}
}

type Claims struct {
	Username string `bson:"username,omitempty"`
	jwt.StandardClaims
}

func (service *AuthenticationService) GenerateToken(userName string) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(service.jwtKey)
	if err != nil {
		panic(err)
	}

	service.JwtToken = tokenString
	service.ExpirationTime = expirationTime
}
