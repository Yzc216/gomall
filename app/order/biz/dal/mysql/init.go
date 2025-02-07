package mysql

import (
	"fmt"
	"github.com/Yzc216/gomall/app/order/biz/model"
	"github.com/Yzc216/gomall/app/order/conf"
	"github.com/Yzc216/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
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
	if err != nil {
		panic(err)
	}
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		if err = DB.AutoMigrate(&model.Order{}, &model.OrderItem{}); err != nil {
			klog.Error("auto migrate order err: ", err)
		}
	}

}
