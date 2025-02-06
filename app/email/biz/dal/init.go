package dal

import (
	"github.com/Yzc216/gomall/app/email/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
