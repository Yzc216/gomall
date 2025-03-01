package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestUpdateCategory_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewUpdateCategoryService(ctx)
	// init req and assert value

	req := &product.UpdateCategoryReq{
		Id: 3,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

//	type CategoryRepoTestSuite struct {
//		suite.Suite
//		repo *model.CategoryRepo
//	}
//
// // 初始化测试套件
//
//	func (suite *CategoryRepoTestSuite) SetupTest() {
//		godotenv.Load("../../.env")
//		dal.Init()
//		suite.repo = model.NewCategoryRepo(mysql.DB)
//	}
//
// /* 测试用例数据准备 */
// var (
//
//	existingCategory = &model.Category{
//		ID:          3,
//		Name:        "智能手机",
//		Description: "高端智能手机分类",
//		ParentID:    2,
//		Level:       2,
//		Sort:        1,
//	}
//
// )
//
// /* 测试用例集合 */
//
//	func (s *CategoryRepoTestSuite) TestUpdateScenarios() {
//		testCases := []struct {
//			name       string
//			req        *product.UpdateCategoryReq
//			wantErr    error
//			verifyFunc func(t *testing.T, updated *model.Category)
//		}{
//			// 测试用例1: 正常多字段更新
//			{
//				name: "valid multi-field update",
//				req: &product.UpdateCategoryReq{
//					Id:          3,
//					Name:        strPtr("旗舰手机"),
//					Description: strPtr("高端旗舰智能手机"),
//					Sort:        int32Ptr(2),
//					ImageUrl:    strPtr("https://example.com/flagship.png"),
//				},
//				verifyFunc: func(t *testing.T, updated *model.Category) {
//					assert.Equal(t, "旗舰手机", updated.Name)
//					assert.Equal(t, "高端旗舰智能手机", updated.Description)
//					assert.EqualValues(t, 2, updated.Sort)
//					assert.Equal(t, "https://example.com/flagship.png", updated.Image)
//
//					// 验证父级未变化
//					assert.EqualValues(t, 2, updated.ParentID)
//					assert.EqualValues(t, 2, updated.Level)
//				},
//			},
//
//			// 测试用例2: 名称冲突
//			{
//				name: "duplicate name in same parent",
//				req: &product.UpdateCategoryReq{
//					Id:   5,
//					Name: strPtr("智能手机"), // 与ID=3冲突
//				},
//				wantErr: types.ErrCategoryNameExists,
//				verifyFunc: func(t *testing.T, _ *model.Category) {
//					// 验证原数据未改变
//					original, err := s.repo.GetByID(context.Background(), 5)
//					assert.NoError(t, err)
//					assert.Equal(t, "智能家居-2024", original.Name)
//				},
//			},
//
//			// 测试用例3: 超长名称
//			//{
//			//	name: "overlength name",
//			//	req: &category.UpdateCategoryReq{
//			//		Id:   6,
//			//		Name: strPtr(strings.Repeat("A", 120)),
//			//	},
//			//	wantErr: category.ErrInvalidUpdate,
//			//},
//			//
//			//// 测试用例5: 非法排序值
//			//{
//			//	name: "negative sort value",
//			//	req: &category.UpdateCategoryReq{
//			//		Id:   2,
//			//		Sort: int32Ptr(-1),
//			//	},
//			//	wantErr: category.ErrInvalidUpdate,
//			//},
//			//
//			//// 测试用例6: 不存在的分类
//			//{
//			//	name: "non-existent category",
//			//	req: &category.UpdateCategoryReq{
//			//		Id:   999,
//			//		Name: strPtr("测试分类"),
//			//	},
//			//	wantErr: category.ErrCategoryNotFound,
//			//},
//			//
//			//// 测试用例7: 仅更新图片
//			//{
//			//	name: "update image only",
//			//	req: &category.UpdateCategoryReq{
//			//		Id:       2,
//			//		ImageUrl: strPtr("https://example.com/new-banner.jpg"),
//			//	},
//			//	verifyFunc: func(t *testing.T, updated *category.Category) {
//			//		assert.Equal(t, "电子产品", updated.Name) // 名称未变
//			//		assert.Equal(t, "https://example.com/new-banner.jpg", updated.Image)
//			//	},
//			//},
//			//
//			//// 测试用例8: 空操作更新
//			//{
//			//	name:    "empty update",
//			//	req:     &category.UpdateCategoryReq{Id: 3},
//			//	wantErr: category.ErrInvalidUpdate,
//			//},
//			//
//			//// 测试用例9: 清空描述
//			//{
//			//	name: "clear description",
//			//	req: &category.UpdateCategoryReq{
//			//		Id:          5,
//			//		Description: strPtr(""), // 显式清空
//			//	},
//			//	verifyFunc: func(t *testing.T, updated *category.Category) {
//			//		assert.Empty(t, updated.Description)
//			//	},
//			//},
//		}
//
//		for _, tc := range testCases {
//			s.T().Run(tc.name, func(t *testing.T) {
//				// 执行更新操作
//				serv := NewUpdateCategoryService(context.Background())
//				_, err := serv.Run(tc.req)
//				if err != nil {
//					return
//				}
//
//				// 错误断言
//				if tc.wantErr != nil {
//					assert.ErrorIs(t, err, tc.wantErr)
//				} else {
//					assert.NoError(t, err)
//				}
//
//				// 结果验证
//				if tc.verifyFunc != nil {
//					updated, err := s.repo.GetByID(context.Background(), tc.req.Id)
//					if assert.NoError(t, err) {
//						tc.verifyFunc(t, updated)
//					}
//				}
//			})
//		}
//	}
//
// /* 辅助函数 */
func strPtr(s string) *string { return &s }
func int32Ptr(i int32) *int32 { return &i }
