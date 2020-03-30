package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "47.107.49.152:6379")
	if err != nil {
		fmt.Println("conn redis failen,", err)
		return
	}
	defer c.Close()
}
