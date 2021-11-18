package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"strings"

	"github.com/luenci/oauth2/store"

	"github.com/google/uuid"
)

func GenerateAuthorizationCode(ctx context.Context, ClientID string) (string, error) {

	code := md5.Sum([]byte(ClientID))
	redisConn, err := store.GetRedisConn()
	if err != nil {
		return "", err
	}
	_, err = redisConn.Do("set", ClientID, string(code[:]))
	if err != nil {
		return "", err
	}

	return string(code[:]), nil
}

func GenerateAuthorizationToken(ctx context.Context, ClientID string) (string, error) {
	buf := bytes.NewBufferString(ClientID)
	buf.WriteString(ClientID)
	token := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	code := base64.URLEncoding.EncodeToString([]byte(token.String()))
	code = strings.ToUpper(strings.TrimRight(code, "="))

	return code, nil
}
