package db

import (
	"encoding/json"
	"physiobot/config"
	"physiobot/modules/common/logger"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

// InitRedis client
func init() {
	configuration := config.GetConfig()

	client = redis.NewClient(&redis.Options{
		Addr:     configuration.Redis.Host + ":" + configuration.Redis.Port,
		Password: configuration.Redis.Password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		logger.Warning(err)
	} else if pong == "PONG" {
		logger.Success("Conected to redis")
	}
}

// Cache get redis instance
func Cache() *redis.Client {
	return client
}

func GetValueFromKey(redisKey string, object interface{}) (error) {
	cache := Cache()
	redisResult, redisError := cache.Get(redisKey).Result()
	if redisError == nil {
		err := json.Unmarshal([]byte(redisResult), object)
		if err == nil {
			return nil
		}
		return err
	}
	return redisError
}

func SetValueFromKey(redisKey string, object interface{}) (error) {
	cache := Cache()
	configuration := config.GetConfig()
	value, err := json.Marshal(object)
	if err == nil {
		return cache.Set(redisKey, value, time.Duration(configuration.Redis.ExpirationTime)*time.Second).Err()
	}
	return err
}