package redis

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestRedisStore_SetValue(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"RedisSetValue", fields{}, args{"key", "value"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PoolInitRedis("127.0.0.1:6379", "123456")
			r := &RedisStore{
				pool: redisPool,
			}
			if err := r.SetValue(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SetValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
