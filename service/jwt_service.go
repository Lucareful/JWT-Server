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

// NewJWTAuthService returns a new JWT service.
func NewJWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    getName(),
	}
}

func getSecretKey() string {
	secret := config.GetConf().JWT.Secret
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func getName() string {
	name := config.GetConf().JWT.Name
	if name == "" {
		name = "Luenci"
	}
	return name
}

func (service *jwtServices) GenerateToken(userId string, isUser bool) string {
	claims := &authCustomClaims{
		userId,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
