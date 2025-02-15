package dal

import (
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
