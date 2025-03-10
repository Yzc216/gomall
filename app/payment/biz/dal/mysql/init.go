package mysql

import (
	"fmt"
	"github.com/Yzc216/gomall/app/payment/biz/model"
	"github.com/Yzc216/gomall/app/payment/conf"
	"github.com/Yzc216/gomall/common/mtl"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.PaymentLog{})
	if err != nil {
		panic(err)
	}
}
