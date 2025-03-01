package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestCreateCategory_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()

	ctx := context.Background()
	s := NewCreateCategoryService(ctx)
	// init req and assert value

	//场景1：创建根分类（无父级）
	//req1 := &product.CreateCategoryReq{
	//	Name:        "数码产品",
	//	Description: "所有数码设备分类",
	//	ParentId:    0, // 0表示根分类
	//	// ImageUrl 不设置
	//}

	// 场景2：创建带图片的子分类
	//req2 := &product.CreateCategoryReq{
	//	Name:        "智能手机",
	//	Description: "高端智能手机分类",
	//	ParentId:    1, // 假设父分类ID=1
	//	ImageUrl:    proto.String("https://cdn.example.com/smartphones.png"),
	//}

	// 场景3：创建无描述的叶子分类

	req := []*product.CreateCategoryReq{
		//{ //2
		//	Name:        "智能手机",
		//	Description: "智能手机分类",
		//	ParentId:    1, // 假设父分类ID=1
		//	ImageUrl:    proto.String("https://cdn.example.com/smartphones.png"),
		//},
		//{ //3
		//	Name:        "高端手机",
		//	Description: "高端智能手机分类",
		//	ParentId:    2, // 假设父分类ID=1
		//	ImageUrl:    proto.String("https://cdn.example.com/smartphones.png"),
		//},
		//{ //4
		//	Name:        "耳机",
		//	Description: "耳机分类",
		//	ParentId:    1, // 假设父分类ID=1
		//	ImageUrl:    proto.String("https://cdn.example.com/smartphones.png"),
		//},
		{ //6
			Name:        "无线耳机",
			Description: "无线耳机分类",
			ParentId:    4, // 假设父分类ID=2
		},
		//{ //5
		//	Name:        "智能家居",
		//	Description: "智能家居设备",
		//	ParentId:    0,
		//},
		{ //7
			Name:        "智能家居-2024",
			Description: "最新智能家居设备",
			ParentId:    5,
			ImageUrl:    proto.String("https://cdn.example.com/smarthome.jpg"),
		},
	}

	for i := 0; i < len(req); i++ {
		resp, err := s.Run(req[i])
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}

}
