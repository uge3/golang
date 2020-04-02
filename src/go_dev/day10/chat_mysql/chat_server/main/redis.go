package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//redis 连接
func initRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     idleConn,    //连接池
		MaxActive:   maxConn,     //连接大小
		IdleTimeout: idleTimeout, //过期时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	return
}

/*
func initRedis(addr string, password string, idleConn, maxConn int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     idleConn,    //连接池
		MaxActive:   maxConn,     //连接大小
		IdleTimeout: idleTimeout, //过期时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c
			//return redis.Dial("tcp", addr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },
	}
	return
}
*/
//取信息
func GetConn() redis.Conn {
	return pool.Get()
}

//关闭redis连接
func PutConn(conn redis.Conn) {
	conn.Close()
}
