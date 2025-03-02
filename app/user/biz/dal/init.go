package dal

import (
	"github.com/7qing/gomall/app/user/biz/dal/mysql"
	"github.com/7qing/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
