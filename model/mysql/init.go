package mysql

import (
	"fmt"
	"github.com/dinzhen12306/template/config"
	"github.com/dinzhen12306/template/golbe"
	"gopkg.in/yaml.v2"
	"log"
	"xorm.io/xorm"
)

type Mysql struct {
	Mysql struct {
		Username  string `yaml:"Username"`
		Password  string `yaml:"Password"`
		Databases string `yaml:"Databases"`
	} `yaml:"Mysql"`
}

// mysql初始化
func Init() error {
	var m Mysql
	str, err := config.GetConf()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(str), &m)
	if err != nil {
		return err
	}
	log.Println(m)
	golbe.XDB, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4", m.Mysql.Username, m.Mysql.Password, m.Mysql.Databases))
	if err != nil {
		return err
	}
	log.Println("mysql连接成功")
	return nil
}

func MigrationDatabases() {
	err := golbe.XDB.Sync2(new(User))
	if err != nil {
		log.Println("数据迁移失败", err)
	}
}
