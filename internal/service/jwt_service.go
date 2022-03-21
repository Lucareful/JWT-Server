package service

import (
	"context"
	"fmt"
	"time"

	"github.com/luenci/oauth2/internal/repository"

	"github.com/luenci/oauth2/config"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	ID   string `json:"id"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	userRepo  repository.UserRepositoryInterface
	secretKey string
	issure    string
}

var _ JWTService = (*jwtServices)(nil)

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

func (srv *jwtServices) GenerateToken(userName, password string, isUser bool) string {
	usr, err := srv.userRepo.GetUserID(context.Background(), userName, password)
	if err != nil {
		return ""
	}
	claims := &authCustomClaims{
		usr.UserId,
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

func NewJWTServices(repository repository.UserRepositoryInterface) *jwtServices {
	return &jwtServices{
		userRepo:  repository,
		secretKey: getSecretKey(),
		issure:    getName(),
	}
}
