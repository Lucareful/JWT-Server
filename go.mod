module github.com/luenci/oauth2

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/google/uuid v1.1.2
	github.com/spf13/viper v1.9.0
	github.com/tidwall/buntdb v1.2.7
)

replace (
	github.com/luenci/oauth2/config => ./config
	github.com/luenci/oauth2/routers => ./routers
	github.com/luenci/oauth2/routers/api => ./api
	github.com/luenci/oauth2/service => ./service
)
