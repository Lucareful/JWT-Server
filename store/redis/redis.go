package store

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool *redis.Pool
)

// PoolInitRedis Redis 连接池.
func PoolInitRedis(server string, password string) {
	// 使 RedisPool 内存逃逸
	redisPool = &redis.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				log.Printf("redis.Dial err:%s", err)
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					log.Printf("redis.Dial err:%s", err)
					c.Close()
					return nil, err
				}
			}
			log.Printf("redis.Dial err:%s", err)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// RedisStore redis存储对象.
type RedisStore struct {
	pool *redis.Pool
}

// GetConn 获取redis连接.
func (r *RedisStore) GetConn() (redis.Conn, error) {
	conn := r.pool.Get()
	return conn, conn.Err()
}

// SetValue 设置redis指定key值.
func (r *RedisStore) SetValue(key, value interface{}) error {
	conn, err := r.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

// GetValue 获取redis指定key值.
func (r *RedisStore) GetValue(key interface{}) (interface{}, error) {
	conn, err := r.GetConn()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return value, nil
}

// DelValue 删除redis指定key值.
func (r *RedisStore) DelValue(key interface{}) error {
	conn, err := r.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

func NewRedisStore() *RedisStore {
	return &RedisStore{redisPool}
}
