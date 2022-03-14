package service

import (
	"fmt"
	"time"

	"github.com/luenci/oauth2/config"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	ID   string `json:"id"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func getSecretKey() string {
	conf := config.GetConf()
	secret := conf.JWT.Secret
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func getName() string {
	conf := config.GetConf()
	name := conf.JWT.Name
	if name == "" {
		name = "Luenci"
	}
	return name
}

func (srv *jwtServices) GenerateToken(userId string, isUser bool) string {
	claims := &authCustomClaims{
		userId,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    getName(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// encoded string
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func (srv *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid JWT token %s", token.Header["alg"])

		}
		return []byte(getSecretKey()), nil
	})

}