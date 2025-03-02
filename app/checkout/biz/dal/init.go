package dal

import (
	"github.com/7qing/gomall/app/checkout/biz/dal/mysql"
	"github.com/7qing/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
