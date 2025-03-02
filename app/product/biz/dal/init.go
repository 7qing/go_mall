package dal

import (
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
