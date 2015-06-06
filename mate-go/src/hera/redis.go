package hera

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisSvc struct {
	connType string
	ip       string
	port     string
	db       string
}

var Redis *RedisSvc = nil
var redisConnection *redis.Pool = nil

func NewRedisSvc() {
	Redis = &RedisSvc{
		connType: "tcp",
		ip:       "192.168.1.182",
		port:     "6379",
		db:       "1",
	}
}

func (this *RedisSvc) getRedisPool() redis.Conn {
	if redisConnection == nil {
		redisConnection = &redis.Pool{
			MaxIdle:     1000,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", this.ip, this.port))
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}

	}
	return redisConnection.Get()
}

func String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

func (this *RedisSvc) Do(cmd string, args ...interface{}) (interface{}, error) {
	enablePool := true
	var c redis.Conn
	var err error
	if enablePool {
		c = this.getRedisPool()
	} else {
		c, err = redis.Dial(this.connType, fmt.Sprintf("%s:%s", this.ip, this.port))
		if err != nil {
			panic(err)
		}
	}
	defer c.Close()
	_, err = c.Do("SELECT", this.db)
	if err != nil {
		panic(err)
	}
	re, err := c.Do(cmd, args...)
	if err != nil {
		panic(err)
	}
	return re, err
}
