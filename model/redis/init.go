package redis

import (
	"github.com/dinzhen12306/template/golbe"
	"github.com/go-redis/redis/v7"
	"log"
)

type Redis struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

// redis初始化
func (r Redis) Init() error {
	golbe.RDB = redis.NewClient(&redis.Options{
		Addr: r.Host + ":" + r.Port,
	})
	if _, err := golbe.RDB.Ping().Result(); err != nil {
		panic(err)
	}
	log.Println("redis连接失败")
	return nil
}
