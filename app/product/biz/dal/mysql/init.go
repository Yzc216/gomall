package mysql

import (
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/conf"
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
	if err != nil {
		panic(err)
	}
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(
		&model.SPU{},
		&model.SKU{},
		&model.Category{},
	)
	if err != nil {
		return
	}

	//categories := []*model.Category{
	//	{Name: "T-Shirt", Description: "T-Shirt"},
	//	{Name: "Sticker", Description: "Sticker"},
	//}
	//DB.CreateInBatches(categories, 2)
	//products := []*model.Product{
	//	{
	//		Name:        "Notebook",
	//		Description: "The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs.",
	//		Picture:     "/static/image/notebook.jpeg",
	//		Price:       9.90,
	//	},
	//	{
	//		Name:        "Mouse-Pad",
	//		Description: "The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience.",
	//		Picture:     "/static/image/mouse-pad.jpeg",
	//		Price:       8.80,
	//	},
	//	{
	//		Name:        "T-Shirt",
	//		Description: "The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.",
	//		Picture:     "/static/image/t-shirt.jpeg",
	//		Price:       6.60,
	//	},
	//	{
	//		Name:        "T-Shirt",
	//		Description: "The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.",
	//		Picture:     "/static/image/t-shirt-1.jpeg",
	//		Price:       2.20,
	//	},
	//	{
	//		Name:        "Sweatshirt",
	//		Description: "The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.",
	//		Picture:     "/static/image/sweatshirt.jpeg",
	//		Price:       1.10,
	//	},
	//	{
	//		Name:        "T-Shirt",
	//		Description: "The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.",
	//		Picture:     "/static/image/t-shirt-2.jpeg",
	//		Price:       1.80,
	//	},
	//	{
	//		Name:        "mascot",
	//		Description: "The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.",
	//		Picture:     "/static/image/logo.jpg",
	//		Price:       4.80,
	//	},
	//}
	//DB.CreateInBatches(products, len(products))
	//DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")

	//if os.Getenv("GO_ENV") != "online" {
	//	//needDemoData := !DB.Migrator().HasTable(&model.Product{})
	//	needDemoData := true
	//	DB.AutoMigrate( //nolint:errcheck
	//		&model.Product{},
	//		&model.Category{},
	//	)
	//	if needDemoData {
	//		DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','','Sticker','Sticker')")
	//		DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', '','Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', '','Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20','', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20','', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', '','Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', '','T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', '','mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
	//		DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
	//	}
	//}
}
