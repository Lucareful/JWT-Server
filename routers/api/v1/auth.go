package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luenci/oauth2/types"
)

func Token(ctx *gin.Context) {

}

func Authorize(ctx *gin.Context) {
	var query types.Authorization
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(query)
	code, err := srv.Authorization.GenerateAuthorizationCode(ctx, query.ClientID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(code)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "success",
	})

}
