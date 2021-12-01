package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"

	"github.com/luenci/oauth2/store"

	"github.com/google/uuid"
)

type authorizationService struct{}

// GenerateAuthorizationCode 生成授权码.
func (a *authorizationService) GenerateAuthorizationCode(ctx context.Context, ClientID string) (int, error) {

	// TODO: 判断客户端是否存在.

	code := rand.Intn(9999999999)
	fmt.Println(code)
	redisStore := store.NewRedisStore()
	if err := redisStore.SetValue(ClientID, code); err != nil {
		return 0, err
	}

	return code, nil
}

// GenerateAccessToken 生成 accessToken.
func (a *authorizationService) GenerateAccessToken(ctx context.Context, AuthCode string) (string, error) {

	buf := bytes.NewBufferString(AuthCode)
	buf.WriteString(AuthCode)
	token := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	code := base64.URLEncoding.EncodeToString([]byte(token.String()))
	code = strings.ToUpper(strings.TrimRight(code, "="))
	redisStore := store.NewRedisStore()
	if err := redisStore.SetValue(AuthCode, code); err != nil {
		return "", err
	}

	return code, nil
}
