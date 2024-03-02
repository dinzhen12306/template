package golbe

import (
	"github.com/go-redis/redis/v7"
	"xorm.io/xorm"
)

var (
	XDB *xorm.Engine
	RDB *redis.Client
)
