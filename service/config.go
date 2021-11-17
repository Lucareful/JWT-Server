package service

import (
	"net/http"
	"time"

	"github.com/luenci/oauth2/pkg"
)

// Config configuration parameters
type Config struct {
	TokenType                   string             // token type
	AllowGetAccessRequest       bool               // to allow GET requests for the token
	AllowedResponseTypes        []pkg.ResponseType // allow the authorization type
	AllowedGrantTypes           []pkg.GrantType    // allow the grant type
	AllowedCodeChallengeMethods []pkg.CodeChallengeMethod
	ForcePKCE                   bool
}

// NewConfig create to configuration instance
func NewConfig() *Config {
	return &Config{
		TokenType:            "Bearer",
		AllowedResponseTypes: []pkg.ResponseType{pkg.Code, pkg.Token},
		AllowedGrantTypes: []pkg.GrantType{
			pkg.AuthorizationCode,
			pkg.PasswordCredentials,
			pkg.ClientCredentials,
			pkg.Refreshing,
		},
		AllowedCodeChallengeMethods: []pkg.CodeChallengeMethod{
			pkg.CodeChallengePlain,
			pkg.CodeChallengeS256,
		},
	}
}

// AuthorizeRequest authorization request
type AuthorizeRequest struct {
	ResponseType        pkg.ResponseType
	ClientID            string
	Scope               string
	RedirectURI         string
	State               string
	UserID              string
	CodeChallenge       string
	CodeChallengeMethod pkg.CodeChallengeMethod
	AccessTokenExp      time.Duration
	Request             *http.Request
}
