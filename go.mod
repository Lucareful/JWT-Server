module github.com/luenci/oauth2

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-session/session v3.1.2+incompatible
	github.com/google/uuid v1.1.2
	github.com/smartystreets/goconvey v1.7.2 // indirect
	github.com/spf13/viper v1.9.0
	github.com/tidwall/buntdb v1.2.7
)

replace (
	github.com/luenci/oauth2/config => ./config
	github.com/luenci/oauth2/pkg => ./pkg
	github.com/luenci/oauth2/routers => ./routers
	github.com/luenci/oauth2/routers/api => ./api
	github.com/luenci/oauth2/service => ./service
)
