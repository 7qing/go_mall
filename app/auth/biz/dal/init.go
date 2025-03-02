package dal

import (
	"github.com/7qing/gomall/app/auth/biz/dal/mysql"
	"github.com/7qing/gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
