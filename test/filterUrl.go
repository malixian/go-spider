package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	_, err = c.Do("sadd","spider","test.com")
	if err != nil{
		log.Printf("redis add url err %v", err)
	}else{
		log.Printf("success")
	}
}
