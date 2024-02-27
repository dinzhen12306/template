package mysql

import (
	nacos "github.com/dinzhen12306/framework/config"
	"log"
)

func init() {
	err := nacos.XDB.Sync(new(User))
	if err != nil {
		log.Println("数据迁移失败")
	}
}
