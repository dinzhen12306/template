package config

import nacos "github.com/dinzhen12306/framework/config"

func NewNacosConfig(fun func() error) error {
	return nacos.Initialisation(&nacos.Nacos{
		IpAddr: "127.0.0.1",
		Port:   8848,
		DataId: "users",
		Group:  "DEFAULT_GROUP",
	}, fun)
}

func GetConf() (string, error) {
	config, err := nacos.GetConfig("users", "DEFAULT_GROUP")
	if err != nil {
		return "", err
	}
	return config, nil
}
