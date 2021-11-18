package config

import (
	"testing"
)

func TestInitConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testRes *Config
			InitConf()
			conf := GetConf()
			if conf == testRes {
				t.Errorf("配置加载失败：%#v", conf)
			}
		})
	}
}
