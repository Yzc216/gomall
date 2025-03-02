package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	"github.com/Yzc216/gomall/app/product/infra/mq"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestCreateProduct_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	mq.Init()
	ctx := context.Background()
	s := NewCreateProductService(ctx)
	// init req and assert value

	resp, err := s.Run(req4)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

var (
	req1 = &product.CreateProductReq{
		BasicInfo: &product.SPUBasicInfo{
			Title:       "Apple iPhone 15 Pro 5G手机",
			SubTitle:    "钛金属设计 A17 Pro芯片 双卡双待",
			Description: "<p>旗舰智能手机，6.1英寸超视网膜XDR显示屏</p>",
			ShopId:      1001,
			Brand:       "Apple",
			Status:      1,
		},
		Media: &product.SPUMedia{
			MainImages: []string{
				"https://cdn.example.com/iphone15_1.jpg",
				"https://cdn.example.com/iphone15_2.jpg",
			},
			VideoUrl: "https://cdn.example.com/iphone15_video.mp4",
		},
		CategoryRelation: &product.CategoryRelation{
			CategoryIds: []uint64{1, 2, 3}, // 手机 > 智能手机 > 高端智能手机
		},
		SKUs: []*product.CreateProductReq_SKUData{
			{
				Title: "iPhone 15 Pro 128GB 黑色钛金属",
				Price: 7999.00,
				Stock: 100,
				Specs: map[string]string{
					"颜色":   "黑色钛金属",
					"存储容量": "128GB",
				},
			},
			{
				Title: "iPhone 15 Pro 256GB 原色钛金属",
				Price: 8999.00,
				Stock: 50,
				Specs: map[string]string{
					"颜色":   "原色钛金属",
					"存储容量": "256GB",
				},
			},
		},
	}

	req2 = &product.CreateProductReq{
		BasicInfo: &product.SPUBasicInfo{
			Title:       "Huawei Mate 60 Pro",
			SubTitle:    "麒麟9000S芯片 卫星通话",
			Description: "<p>6.82英寸OLED曲面屏，IP68防水</p>",
			ShopId:      1002,
			Brand:       "Huawei",
			Status:      product.SPUStatus_PENDING_REVIEW,
		},
		Media: &product.SPUMedia{
			MainImages: []string{"https://cdn.example.com/mate60_1.jpg"},
			VideoUrl:   "https://cdn.example.com/mate60_video.mp4",
		},
		CategoryRelation: &product.CategoryRelation{
			CategoryIds: []uint64{1, 2, 3}, // 数码产品→智能手机→高端手机
		},
		SKUs: []*product.CreateProductReq_SKUData{
			{
				Title: "Mate 60 Pro 12+512GB 雅川青",
				Price: 6999.00,
				Stock: 50,
				Specs: map[string]string{
					"颜色":   "雅川青",
					"内存组合": "12GB+512GB",
				},
			},
		},
	}

	req3 = &product.CreateProductReq{
		BasicInfo: &product.SPUBasicInfo{
			Title:  "智能温控器套装",
			ShopId: 4005,
			Brand:  "Honeywell",
			Status: product.SPUStatus_APPROVED,
		},
		Media: &product.SPUMedia{
			MainImages: []string{"https://cdn.example.com/mate60_1.jpg"},
			VideoUrl:   "https://cdn.example.com/mate60_video.mp4",
		},
		CategoryRelation: &product.CategoryRelation{
			CategoryIds: []uint64{5, 7}, // 智能家居→智能家居-2024
		},
		SKUs: []*product.CreateProductReq_SKUData{
			{
				Title: "基础版-白色",
				Price: 299.00,
				Specs: map[string]string{
					"型号": "T-100",
					"颜色": "白色",
					"功率": "2200W",
				},
			},
			{
				Title: "Pro版-黑色",
				Price: 599.00,
				Specs: map[string]string{
					"型号":   "T-200",
					"颜色":   "黑色",
					"功率":   "3000W",
					"附加功能": "语音控制",
				},
			},
		},
	}

	req4 = &product.CreateProductReq{
		BasicInfo: &product.SPUBasicInfo{
			Title:       "Sony WH-1000XM5 无线降噪耳机",
			SubTitle:    "AI智能降噪 30小时续航",
			Description: "<p>行业领先的主动降噪技术</p>",
			ShopId:      2003,
			Brand:       "Sony",
			Status:      product.SPUStatus_ONLINE,
		},
		Media: &product.SPUMedia{
			MainImages: []string{"https://cdn.example.com/sony_earphones.jpg"},
		},
		CategoryRelation: &product.CategoryRelation{
			CategoryIds: []uint64{4, 6}, // 耳机→无线耳机
		},
		SKUs: []*product.CreateProductReq_SKUData{
			{
				Title: "铂金银标准版",
				Price: 2499.00,
				Stock: 200,
				Specs: map[string]string{
					"颜色":   "铂金银",
					"续航时间": "30小时",
				},
			},
		},
	}
)
