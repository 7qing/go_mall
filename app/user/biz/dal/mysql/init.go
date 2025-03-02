package mysql

import (
	"fmt"
	"github.com/7qing/gomall/app/user/biz/model"
	"github.com/7qing/gomall/app/user/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 修改mysql的启动服务样例，从配置中启动 (正常时候)
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	// 测试时候
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	// 迁移表
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		panic(err) // 迁移失败时处理错误
	}
	if err != nil {
		panic(err)
	}
}
