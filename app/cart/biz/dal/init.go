package dal

import (
	"github.com/7qing/gomall/app/cart/biz/dal/mysql"
	"github.com/7qing/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
