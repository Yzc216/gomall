package dal

import (
	"github.com/Yzc216/gomall/app/frontend/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
