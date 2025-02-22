package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type AttributeKeyRepoTestSuite struct {
	suite.Suite
	repo *AttributeKeyRepo
}

func (suite *AttributeKeyRepoTestSuite) SetupTest() {
	godotenv.Load("../../.env")
	dal.Init()
	// 使用已初始化的数据库
	suite.repo = NewAttributeKeyRepo(context.Background(), mysql.DB)

	// 清空测试数据
	suite.cleanTestData()
}

func (suite *AttributeKeyRepoTestSuite) TearDownTest() {
	// 测试结束后清理数据
	suite.cleanTestData()
}

func (suite *AttributeKeyRepoTestSuite) cleanTestData() {
	// 使用事务保证清理操作的原子性
	err := suite.repo.db.Transaction(func(tx *gorm.DB) error {
		return tx.Exec("DELETE FROM attribute_key").Error
	})
	suite.NoError(err, "清理测试数据失败")
}

func TestAttributeKeyRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AttributeKeyRepoTestSuite))
}

func (suite *AttributeKeyRepoTestSuite) createTestAttribute(attr *model.AttributeKey) {
	err := suite.repo.Create(attr)
	suite.NoError(err, "创建测试属性失败")
}

// List 方法测试
func (suite *AttributeKeyRepoTestSuite) TestList() {
	// 准备测试数据
	suite.createTestAttribute(&model.AttributeKey{Name: "color", DataType: "string", IsFilter: true, Order: 2})
	suite.createTestAttribute(&model.AttributeKey{Name: "size", DataType: "int", IsFilter: true, Order: 1})
	suite.createTestAttribute(&model.AttributeKey{Name: "weight", DataType: "float", IsFilter: false, Order: 3})

	// 测试分页和排序
	suite.Run("分页和排序", func() {
		p := Pagination{Page: 1, PageSize: 2}
		filter := AttributeKeyFilter{}
		result, err := suite.repo.List(filter, p)

		suite.NoError(err)
		suite.Len(result, 2)
		suite.Equal("size", result[0].Name) // 按order排序
	})

	// 测试名称过滤
	suite.Run("名称过滤", func() {
		p := Pagination{Page: 1, PageSize: 10}
		filter := AttributeKeyFilter{Name: "s"}
		result, err := suite.repo.List(filter, p)

		suite.NoError(err)
		suite.Len(result, 1)
		suite.Equal("size", result[0].Name)
	})

	// 测试数据类型过滤
	suite.Run("数据类型过滤", func() {
		p := Pagination{Page: 1, PageSize: 10}
		filter := AttributeKeyFilter{DataType: "string"}
		result, err := suite.repo.List(filter, p)

		suite.NoError(err)
		suite.Len(result, 1)
		suite.Equal("color", result[0].Name)
	})

	// 测试IsFilter过滤
	suite.Run("IsFilter过滤", func() {
		var falseCount int64
		suite.repo.db.Model(&model.AttributeKey{}).Where("is_filter = ?", 0).Count(&falseCount)
		suite.Equal(int64(1), falseCount, "测试数据中应有1条is_filter=false的记录")
		isFilter := true
		p := Pagination{Page: 1, PageSize: 10}
		filter := AttributeKeyFilter{IsFilter: &isFilter}
		result, err := suite.repo.List(filter, p)

		suite.NoError(err)
		suite.Len(result, 2)
	})
}

// Count 方法测试
func (suite *AttributeKeyRepoTestSuite) TestCount() {
	suite.createTestAttribute(&model.AttributeKey{Name: "test1", DataType: "string"})
	suite.createTestAttribute(&model.AttributeKey{Name: "test2", DataType: "int"})

	suite.Run("无过滤条件", func() {
		count, err := suite.repo.Count(AttributeKeyFilter{})
		suite.NoError(err)
		suite.Equal(int64(2), count)
	})

	suite.Run("名称过滤", func() {
		count, err := suite.repo.Count(AttributeKeyFilter{Name: "test1"})
		suite.NoError(err)
		suite.Equal(int64(1), count)
	})
}

// GetByID 方法测试
func (suite *AttributeKeyRepoTestSuite) TestGetByID() {
	attr := &model.AttributeKey{Name: "test"}
	suite.createTestAttribute(attr)

	suite.Run("存在记录", func() {
		result, err := suite.repo.GetByID(attr.KeyID)
		suite.NoError(err)
		suite.Equal(attr.Name, result.Name)
	})

	suite.Run("不存在记录", func() {
		_, err := suite.repo.GetByID(999)
		suite.Error(err)
		suite.True(errors.Is(gorm.ErrRecordNotFound, err))
	})
}

// Create 方法测试
func (suite *AttributeKeyRepoTestSuite) TestCreate() {
	suite.Run("创建成功", func() {
		attr := &model.AttributeKey{Name: "unique"}
		err := suite.repo.Create(attr)
		suite.NoError(err)
		suite.NotZero(attr.KeyID)
	})

	suite.Run("名称冲突", func() {
		attr := &model.AttributeKey{Name: "duplicate"}
		suite.NoError(suite.repo.Create(attr))

		err := suite.repo.Create(&model.AttributeKey{Name: "duplicate"})
		suite.Error(err)
		suite.Contains(err.Error(), "already exists")
	})
}

// Update 方法测试
func (suite *AttributeKeyRepoTestSuite) TestUpdate() {
	attr := &model.AttributeKey{Name: "original"}
	suite.createTestAttribute(attr)

	suite.Run("正常更新", func() {
		attr.Name = "updated"
		err := suite.repo.Update(attr)
		suite.NoError(err)

		updated, err := suite.repo.GetByID(attr.KeyID)
		suite.NoError(err)
		suite.Equal("updated", updated.Name)
	})

	suite.Run("名称冲突", func() {
		otherAttr := &model.AttributeKey{Name: "other"}
		suite.createTestAttribute(otherAttr)

		attr.Name = "other"
		err := suite.repo.Update(attr)
		suite.Error(err)
		suite.Contains(err.Error(), "name conflict")
	})
}

// Delete 方法测试
func (suite *AttributeKeyRepoTestSuite) TestDelete() {
	attr := &model.AttributeKey{Name: "to-delete"}
	suite.createTestAttribute(attr)

	suite.Run("无关联值", func() {
		err := suite.repo.Delete(attr.KeyID)
		suite.NoError(err)

		_, err = suite.repo.GetByID(attr.KeyID)
		suite.Error(err)
	})

	suite.Run("有关联值", func() {
		// 创建关联值
		suite.repo.db.Model(&model.AttributeValue{}).Create(&model.AttributeValue{KeyID: attr.KeyID})

		err := suite.repo.Delete(attr.KeyID)
		suite.Error(err)
		suite.Contains(err.Error(), "related values")
	})
}

func (suite *AttributeKeyRepoTestSuite) BenchmarkListQuery() {
	// 准备大量测试数据
	for i := 0; i < 1000; i++ {
		attr := &model.AttributeKey{
			Name:     fmt.Sprintf("perf_test_%d", i),
			DataType: "string",
			IsFilter: true,
			Order:    i,
		}
		suite.createTestAttribute(attr)
	}

	suite.Run("基准测试", func() {
		result, err := suite.repo.List(
			AttributeKeyFilter{Name: "perf_test"},
			Pagination{Page: 1, PageSize: 100},
		)
		suite.NoError(err)
		suite.Len(result, 100)
	})
}
