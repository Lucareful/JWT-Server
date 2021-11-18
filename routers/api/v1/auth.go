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
		return
	}
	fmt.Println(query)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
