package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationService struct {
	jwtKey []byte

	JwtToken       string
	ExpirationTime time.Time
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{
		jwtKey: []byte("Some random key for the time being"),
	}
}

type Claims struct {
	Username string `bson:"username,omitempty"`
	jwt.StandardClaims
}

func (service *AuthenticationService) GenerateToken(userName string) {
	expirationTime := time.Now().Local().Add(time.Hour*time.Duration(5) +
		time.Minute*time.Duration(5) +
		time.Second*time.Duration(5))

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
