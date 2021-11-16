module github.com/luenci/oauth2

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/spf13/viper v1.9.0
)

replace (
	github.com/luenci/oauth2/config => ./config
	github.com/luenci/oauth2/routers => ./routers
	github.com/luenci/oauth2/routers/api => ./api
)