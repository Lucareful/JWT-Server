module github.com/luenci/oauth2

go 1.16

require (
	github.com/cheggaaa/pb/v3 v3.0.8
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/gomodule/redigo v1.8.5
	github.com/google/uuid v1.1.2
	github.com/luenci/errors v0.0.0-20211126040220-422339e4271f // indirect
	github.com/luenci/gopkg v0.0.0-20211126042024-4242870fb0e6
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/spf13/viper v1.9.0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

replace (
	github.com/luenci/oauth2/config => ./config
	github.com/luenci/oauth2/pkg => ./pkg
	github.com/luenci/oauth2/routers => ./routers
	github.com/luenci/oauth2/routers/api => ./api
	github.com/luenci/oauth2/service => ./service
)
