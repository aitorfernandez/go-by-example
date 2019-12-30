package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Fatal(err)
			}
			return conn, err
		},
		MaxActive: 12000,
		MaxIdle:   80,
	}
}

var pool = NewPool()

// Set key to hold the string value
func Set(key string, value interface{}) error {
	p := pool.Get()
	defer p.Close()

	if _, err := p.Do("SET", key, value); err != nil {
		return err
	}

	return nil
}

// Hset field in the hash stored at key to value
func Hset(key string, field string, value interface{}) error {
	p := pool.Get()
	defer p.Close()

	if _, err := p.Do("HSET", key, field, value); err != nil {
		return err
	}

	return nil
}

// Get the value of key
func Get(key string) string {
	p := pool.Get()
	defer p.Close()

	v, _ := redis.String(p.Do("GET", key))

	return v
}

// Hget returns the value associated with field in the hash stored at key
func Hget(key string, field string) string {
	p := pool.Get()
	defer p.Close()

	v, _ := redis.String(p.Do("HGET", key, field))

	return v
}
