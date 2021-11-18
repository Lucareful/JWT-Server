package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

func GenerateAuthorizationCode(ctx context.Context, ClientID string) (string, error) {

}

func GenerateAuthorizationToken(ctx context.Context, ClientID string) (string, error) {
	buf := bytes.NewBufferString(ClientID)
	buf.WriteString(ClientID)
	token := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	code := base64.URLEncoding.EncodeToString([]byte(token.String()))
	code = strings.ToUpper(strings.TrimRight(code, "="))

	return code, nil
}
