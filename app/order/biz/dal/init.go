package dal

import (
	"github.com/7qing/gomall/app/order/biz/dal/mysql"
	"github.com/7qing/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
