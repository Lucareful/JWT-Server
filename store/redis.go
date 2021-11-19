package store

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisPool 声明全局连接池句柄
var RedisPool *redis.Pool

// GetRedisConn 获取 redis 链接对象
func GetRedisConn() (redis.Conn, error) {
	// 从连接池中获取 连接对象
	conn := RedisPool.Get()

	if conn.Err() != nil {
		return nil, errors.New("redis conn 错误")
	}

	return conn, nil
}

// PoolInitRedis Redis 连接池
func PoolInitRedis(server string, password string, DB int) {
	// 使 RedisPool 内存逃逸
	RedisPool = &redis.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(DB))
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
