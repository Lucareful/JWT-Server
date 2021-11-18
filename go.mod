module github.com/luenci/oauth2

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/gomodule/redigo v1.8.5
	github.com/google/uuid v1.1.2
	github.com/spf13/viper v1.9.0
)

replace (
	github.com/luenci/oauth2/config => ./config
	github.com/luenci/oauth2/pkg => ./pkg
	github.com/luenci/oauth2/routers => ./routers
	github.com/luenci/oauth2/routers/api => ./api
	github.com/luenci/oauth2/service => ./service
)
