package redis

import (
	"github.com/go-redis/redis"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"time"
)

type ClientRedis struct {
	Object *redis.Client
}

var Client *ClientRedis

func new() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
	})
	if client != nil {
		log.Infof("redis client create success")
	}
	return client
}

func (c *ClientRedis) Init() {
	Client = &ClientRedis{Object: new()}
}

func (c *ClientRedis) Close() {
	Client.Object.Close()
}

func (c *ClientRedis) Get(key string) string {
	val, err := Client.Object.Get(key).Result()
	// 判断查询是否出错
	if err != nil {
		return ""
	}
	return val
}

func (c *ClientRedis) Set(key, val string, time time.Duration) error {
	// 第三个参数代表key的过期时间，0代表不会过期。
	err := Client.Object.Set(key, val, time).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientRedis) HSet(key, filed string, val interface{}) error {
	err := Client.Object.HSet(key, filed, val).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientRedis) HGet(key, filed string) interface{} {
	val, err := Client.Object.HGet(key, filed).Result()
	if err != nil {
		log.Errorf(err, "get hash val fail. key:%s, filed: %s", key, filed)
		return nil
	}
	return val
}

func (c *ClientRedis) HGetAll(key string) (map[string]string, error) {
	// 一次性返回key=user_1的所有hash字段和值
	data, err := Client.Object.HGetAll(key).Result()
	return data, err
}

func (c *ClientRedis) HSetAll(key string, data map[string]interface{}) error {
	// 一次性保存多个hash字段值
	err := Client.Object.HMSet(key, data).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientRedis) Expire(key string, duration time.Duration) error {
	err := Client.Object.Expire(key, duration).Err()
	if err != nil {
		return err
	}
	return nil
}
