package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var c redis.Conn
var err error

type Redis struct {
	Name string
}

func init()  {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
}

func (r *Redis) Hredis()  {
	fmt.Println("redis loading...")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	_, err = c.Do("SET", "test", "testxxx")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	v, err := redis.String(c.Do("GET", "test"))
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	fmt.Println("redis test value is:", v)
	defer c.Close()
}

func (r *Redis) Push(key string) {
	fmt.Println("push redis")
}

func (r *Redis) Pull(key string){
	fmt.Println("pull redis")
}
