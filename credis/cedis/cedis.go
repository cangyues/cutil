package cedis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var cedis *redis.Pool

func RedisPoolInit(server string, password string) {
	cedis = &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 240 * time.Second,
		MaxActive:   15,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
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

func GetString(key string) string {
	v, _ := redis.String(redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("GET", key)
	}), nil)
	return v
}

func GetInt(key string) int {
	v, _ := redis.Int(redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("GET", key)
	}), nil)
	return v
}

func GetInt64(key string) int64 {
	v, _ := redis.Int64(redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("GET", key)
	}), nil)
	return v
}

func GetFloat64(key string) float64 {
	v, _ := redis.Float64(redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("GET", key)
	}), nil)
	return v
}

func Del(key string) {
	redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("DEL", key)
	})
}

func Set(key string, value interface{}) {
	redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("SET", key, value)
	})
}

func SetEX(key string, value interface{}, time int) {
	redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("SET", key, value, "EX", time)
	})
}

//检查Key是否过期
func ExistsKey(key string) bool {
	v, _ := redis.Bool(redis.Float64(redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("EXISTS", key)
	}), nil))
	return v
}

//设置key过期时间
func SetKeyEX(key string, time int) {
	redisCMD(func(conn redis.Conn) (i interface{}, err error) {
		return conn.Do("EXPIRE", key, time)
	})
}

func redisCMD(cmd func(conn redis.Conn) (interface{}, error)) interface{} {
	con := cedis.Get()
	defer con.Close()
	v, e := cmd(con)
	if e != nil {
		panic(e)
	}
	return v
}
