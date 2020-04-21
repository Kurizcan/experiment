package redis

import (
	"experiment/config"
	"fmt"
	"testing"
	"time"
)

func TestClientRedis_Set(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}

	Client.Init()
	defer Client.Close()
	Client.Set("lisa", "wonderful", time.Hour*2)
}

func TestClientRedis_Get(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}

	Client.Init()
	defer Client.Close()
	val := Client.Get("lisa")
	fmt.Print(val)
}

func TestClientRedis_HGet(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}

	Client.Init()
	defer Client.Close()
	val := Client.HGet("HashTest", "1")
	fmt.Println(val)
}

func TestClientRedis_HSet(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}

	Client.Init()
	defer Client.Close()
	Client.HSet("HashTest", "1", 1)
	Client.HSet("HashTest", "2", 2)
}

func TestClientRedis_HGetAll(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}

	Client.Init()
	defer Client.Close()
	Client.HGetAll("HashTest")
}
