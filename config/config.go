package config

import nacos "github.com/dinzhen12306/framework/config"

func init() {
	config := nacos.Initialisation(&nacos.Nacos{
		IpAddr: "127.0.0.1",
		Port:   8848,
		DataId: "users",
		Group:  "DEFAULT_GROUP",
	})
	nacos.NewDatabasesConn(config.Mysql)
}
