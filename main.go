package main

import "fmt"
import "github.com/garyburd/redigo/redis"
import (
	"Record"
	"code.google.com/p/goprotobuf/proto"
)

func main() {
	c, err := redis.Dial("tcp", "10.10.21.158:6381")
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := c.Do("LINDEX", "1426008668", 1)
	resStr := string(v.([]uint8))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resStr)
	}
	outMsg := &Record.Record{}
	err = proto.Unmarshal(v.([]uint8), outMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(outMsg)
}
