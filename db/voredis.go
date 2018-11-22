package db

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var RedisClient *redis.Pool


func RedisGo(maxidle int, maxactivity int, idletimeout time.Duration, host string, timeout time.Duration,
	maxwaittime time.Duration, tbermillis time.Duration, passwd string, dbindex int) {
	RedisClient = &redis.Pool{
		MaxIdle: maxidle,
		MaxActive: maxactivity,
		IdleTimeout: idletimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialReadTimeout(timeout),
				redis.DialWriteTimeout(timeout), redis.DialConnectTimeout(maxwaittime),
				redis.DialPassword(passwd))
			if err != nil {
				log.Println(err)
				return nil, err
			}
			c.Do("SELECT", dbindex)

			return c, err

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {

			if time.Since(t) < tbermillis {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
		Wait: true,
	}
}
