package dal

import (
	"github.com/7qing/gomall/app/payment/biz/dal/mysql"
	"github.com/7qing/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
