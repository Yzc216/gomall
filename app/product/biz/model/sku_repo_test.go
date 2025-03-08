package model

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SKURepoTestSuite struct {
	suite.Suite

	repo  *SKURepo
	spuID uint64 // 用于关联测试的SPU ID
}

// 初始化测试套件
func (s *SKURepoTestSuite) SetupSuite() {
	godotenv.Load("../../.env")
	dal.Init()
	s.repo = NewSKURepo(context.Background(), mysql.DB)

	// 确保表结构存在
	s.repo.db.AutoMigrate(&SKU{}, &AttributeValue{}, &AttributeKey{})

	// 创建测试用的SPU
	spu := &SPU{Title: "Test SPU"}
	s.repo.db.Create(spu)
	s.spuID = spu.ID
}

// 每个测试前的清理
func (s *SKURepoTestSuite) SetupTest() {
	// 清空相关表
	s.repo.db.Exec("DELETE FROM sku")
	s.repo.db.Exec("DELETE FROM attribute_value")
	s.repo.db.Exec("DELETE FROM attribute_key")
}

// 测试保存SKU
func (s *SKURepoTestSuite) TestSave() {
	// 准备测试数据
	sku := &SKU{
		ID:    s.spuID,
		Title: "Test SKU",
		Specs: []AttributeValue{
			{KeyID: 1, Value: "Red"},
			{KeyID: 2, Value: "XL"},
		},
	}

	// 执行保存
	err := s.repo.Save(sku)
	s.NoError(err)

	// 验证数据库记录
	var dbSKU SKU
	s.repo.db.Preload("Specs").First(&dbSKU, sku.ID)
	s.Equal(2, len(dbSKU.Specs))
}

// 测试批量保存
func (s *SKURepoTestSuite) TestSaveBatch() {
	skus := make([]*SKU, 2)
	for i := 0; i < 2; i++ {
		skus[i] = &SKU{
			ID:    s.spuID,
			Title: fmt.Sprintf("SKU-%d", i),
			Specs: []AttributeValue{
				{KeyID: 1, Value: "Value"},
			},
		}
	}

	err := s.repo.SaveBatch(skus)
	s.NoError(err)

	var count int64
	s.repo.db.Model(&SKU{}).Count(&count)
	s.Equal(int64(2), count)
}

// 测试查询SKU
func (s *SKURepoTestSuite) TestQuery() {
	// 先创建测试数据
	sku := &SKU{ID: s.spuID, Title: "Query Test"}
	s.repo.db.Create(sku)

	// 执行查询
	result, err := s.repo.Query(sku.ID)
	s.NoError(err)
	s.Equal(sku.Title, result.Title)
}

// 测试更新状态
func (s *SKURepoTestSuite) TestUpdateStatus() {
	sku := &SKU{ID: s.spuID, IsActive: false}
	s.repo.db.Create(sku)

	err := s.repo.UpdateStatusBySkuId(sku.ID, true)
	s.NoError(err)

	var updated SKU
	s.repo.db.First(&updated, sku.ID)
	s.True(updated.IsActive)
}

// 测试删除SKU
func (s *SKURepoTestSuite) TestDelete() {
	sku := &SKU{ID: s.spuID}
	s.repo.db.Create(sku)

	err := s.repo.Delete(sku.ID)
	s.NoError(err)

	var count int64
	s.repo.db.Model(&SKU{}).Where("id = ?", sku.ID).Count(&count)
	s.Equal(int64(0), count)
}

// 测试全量更新属性
func (s *SKURepoTestSuite) TestUpdateSKUAttributes() {
	// 创建属性键
	keys := []AttributeKey{
		{Name: "color"},
		{Name: "size"},
	}
	s.repo.db.Create(&keys)

	// 创建初始SKU
	sku := &SKU{ID: s.spuID}
	s.repo.db.Create(sku)

	// 新属性
	newSpecs := []AttributeValue{
		{KeyID: keys[0].KeyID, Value: "Blue"},
		{KeyID: keys[1].KeyID, Value: "L"},
	}

	err := s.repo.UpdateSKUAttributes(sku.ID, newSpecs)
	s.NoError(err)

	// 验证属性
	var attrs []AttributeValue
	s.repo.db.Where("sku_id = ?", sku.ID).Find(&attrs)
	s.Equal(2, len(attrs))
}

// 测试增量更新属性
func (s *SKURepoTestSuite) TestPatchSKUAttributes() {
	// 创建属性键
	keys := []AttributeKey{
		{Name: "color", KeyID: 1},
		{Name: "size", KeyID: 2},
	}
	s.repo.db.Create(&keys)

	// 创建初始SKU和属性
	sku := &SKU{ID: s.spuID}
	s.repo.db.Create(sku)
	s.repo.db.Create(&AttributeValue{SkuID: sku.ID, KeyID: 1, Value: "Red"})

	// 执行更新
	updates := map[string]string{
		"color": "Black",
		"size":  "XXL",
	}

	err := s.repo.PatchSKUAttributes(s.repo.ctx, sku.ID, updates)
	s.NoError(err)

	// 验证结果
	var attrs []AttributeValue
	s.repo.db.Where("sku_id = ?", sku.ID).Find(&attrs)
	s.Equal(2, len(attrs))
}

// 测试无效属性名称
func (s *SKURepoTestSuite) TestInvalidAttributeNames() {
	updates := map[string]string{
		"invalid_attr": "value",
	}

	err := s.repo.PatchSKUAttributes(s.repo.ctx, 1, updates)
	s.Error(err)
	s.Contains(err.Error(), "无效的属性名称")
}

// 测试并发解析属性名称
func (s *SKURepoTestSuite) TestResolveNamesConcurrently() {
	// 创建测试属性键
	keys := []AttributeKey{
		{Name: "color"},
		{Name: "size"},
	}
	s.repo.db.Create(&keys)

	updates := map[string]string{
		"color": "Red",
		"size":  "XL",
	}

	nameToID, invalid, err := s.repo.resolveAttributeNamesConcurrently(s.repo.db, updates)
	s.NoError(err)
	s.Empty(invalid)
	s.Equal(keys[0].KeyID, nameToID["color"])
}

// 运行测试套件
func TestSKURepoSuite(t *testing.T) {
	suite.Run(t, new(SKURepoTestSuite))
}
