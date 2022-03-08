package v1

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/luenci/gopkg"
	"github.com/luenci/oauth2/types"
)

func Token(ctx *gin.Context) {
	var query types.AccessToken

	if err := ctx.ShouldBindQuery(&query); err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}
	code, err := srv.Authorization.GenerateAccessToken(ctx, query.AuthCode)
	if err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}

	pkg.Response(ctx, 200000, code)

}

// Authorize 生成授权码.
func Authorize(ctx *gin.Context) {
	var query types.Authorization
	if err := ctx.ShouldBindQuery(&query); err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}
	code, err := srv.Authorization.GenerateAuthorizationCode(ctx, query.ClientID)
	if err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}

	pkg.Response(ctx, 200000, code)
}
