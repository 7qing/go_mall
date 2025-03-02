package dal

import (
	"github.com/7qing/gomall/api/frontend/biz/dal/mysql"
	"github.com/7qing/gomall/api/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
