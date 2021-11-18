package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/luenci/oauth2/config"
	v1 "github.com/luenci/oauth2/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter(conf *config.Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(conf.Server.Mode)

	//// programmatically set swagger info
	//docs.SwaggerInfo.Title = "Study Swagger API"
	//docs.SwaggerInfo.Description = "This is a sample server API."
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "127.0.0.1:8000"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//
	//r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/oauth2/authorize", v1.Authorize)

		// 校验 Token
		apiv1.GET("/oauth2/token", v1.Token)
	}

	return r
}
