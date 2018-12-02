package redis

import (
  "fmt"
  "github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() {
  RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
  if err != nil {
  	fmt.Println(err)
  }
}

func StoreKeyValue(key string, value interface{}) error{
  err := RedisClient.Set(key, value, 0).Err()
  if err != nil {
		return err
	}
  return nil
}

func GetKeyValue(key string) (interface{}, error) {
  val, err := RedisClient.Get(key).Result()
	if err != nil {
		return err.Error(), err
	}
  return val, nil
}
