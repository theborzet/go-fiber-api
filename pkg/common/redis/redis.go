package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func SetupClient() {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADD"),
		Password: viper.GetString("REDIS_PASS"),
		DB:       viper.GetInt("REDIS_DB"),
	})
}
