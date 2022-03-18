package v1

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/luenci/gopkg"
)

func Login(ctx *gin.Context) {

	var user struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}

	token := srv.JWT.GenerateToken(users.UserId, true)

	pkg.Response(ctx, 200000, token)

}
