package v1

import (
	"net/http"
	"net/url"

	"github.com/luenci/oauth2/service"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

var (
	srv = service.GetService()
)

func CheckToken(ctx *gin.Context) {
}

func Authorize(ctx *gin.Context) {

	store, err := session.Start(ctx, ctx.Writer, ctx.Request)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	ctx.Request.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = srv.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
	}

}
