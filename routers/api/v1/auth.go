package v1

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/luenci/gopkg"
	"github.com/luenci/oauth2/routers/api/schema"
)

func Token(ctx *gin.Context) {
	var query schema.AccessToken

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(400001, schema.Translate(err))
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
	var query schema.Authorization
	if err := ctx.ShouldBindQuery(&query); err != nil {
		pkg.Response(ctx, 400001, schema.Translate(err))
		return
	}
	code, err := srv.Authorization.GenerateAuthorizationCode(ctx, query.ClientID)
	if err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}

	pkg.Response(ctx, 200000, code)
}
