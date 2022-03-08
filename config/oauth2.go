package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	cfgFile = "./config/oauth2.yaml"
)

var (
	cfg Config
)

type Config struct {
	Server Server
	Mysql  Mysql
	Log    Log
	Oauth2 Oauth2
	Redis  Redis
}

type Oauth2 struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	Endpoint     string
}

type Redis struct {
	Host     string
	Password string
	DB       int
}

type Server struct {
	Mode         string
	BindAddress  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Mysql struct {
	DSN                   string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
}

type Log struct {
	Level string
}

// InitConf 初始化加载配置
func InitConf() {
	viper.SetConfigFile(cfgFile)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}

// GetConf 获取配置信息
func GetConf() *Config {
	return &cfg
}
