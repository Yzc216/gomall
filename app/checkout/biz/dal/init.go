package dal

import (
	"github.com/Yzc216/gomall/app/checkout/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
