package v1

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/luenci/gopkg"
	msRepo "github.com/luenci/oauth2/repository/mysql"
	"github.com/luenci/oauth2/service"
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
	srv := service.NewJWTServices(msRepo.NewUserRepository())
	token := srv.GenerateToken(user.Name, user.Password, true)

	pkg.Response(ctx, 200000, token)

}
