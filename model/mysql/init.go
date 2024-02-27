package mysql

import (
	nacos "github.com/dinzhen12306/framework/config"
	"log"
)

func MigrationDatabases() {
	err := nacos.XDB.Sync2(new(User))
	if err != nil {
		log.Println("数据迁移失败")
	}
}
