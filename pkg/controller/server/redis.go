package server

import (
        "os"
        "github.com/go-redis/redis"
       )

func redis_init() *redis.Client{
  
  client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_CONN"),
  })

  _,err := client.Ping().Result()

  if err != nil {
     panic("Can't talk to Redis")  
  }

  return client
}
