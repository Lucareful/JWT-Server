package v1

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/luenci/gopkg"
	"github.com/luenci/oauth2/models"
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

	users := models.NewUser()
	err := users.GetUserID(user.Name, user.Password)

	ctx.SetCookie("UserId", users.UserId, 3600, "/", "localhost", false, true)

	if err != nil {
		pkg.Response(ctx, 400001, err)
		return
	}

	pkg.Response(ctx, 200000, users)

}
