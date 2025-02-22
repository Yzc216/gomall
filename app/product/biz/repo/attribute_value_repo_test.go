package repo

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type AttributeValueRepoTestSuite struct {
	suite.Suite
	repo *AttributeValueRepo
}

func (suite *AttributeValueRepoTestSuite) SetupTest() {
	godotenv.Load("../../.env")
	dal.Init()
	// 使用已初始化的 MySQL 数据库
	suite.repo = NewAttributeValueRepo(context.Background(), mysql.DB)

	// 清空测试表
	suite.cleanTables()
}

func (suite *AttributeValueRepoTestSuite) TearDownTest() {
	suite.cleanTables()
}

func (suite *AttributeValueRepoTestSuite) cleanTables() {
	// 使用事务保证清理操作的原子性
	err := suite.repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM attribute_value").Error; err != nil {
			return err
		}
		return nil
	})
	suite.NoError(err, "清理测试数据失败")
}

func TestAttributeValueRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AttributeValueRepoTestSuite))
}

// 创建测试数据辅助函数
func (suite *AttributeValueRepoTestSuite) createTestValues(values []model.AttributeValue) {
	err := suite.repo.SaveBatch(values)
	suite.NoError(err, "创建测试数据失败")
}

// --- SaveBatch 测试 ---
func (suite *AttributeValueRepoTestSuite) TestSaveBatch() {
	// 准备测试数据
	testData := []model.AttributeValue{
		{SkuID: 1001, KeyID: 1, Value: "Red"},
		{SkuID: 1001, KeyID: 2, Value: "XL"},
	}

	suite.Run("首次插入新记录", func() {
		err := suite.repo.SaveBatch(testData)
		suite.NoError(err)

		// 验证数据插入
		var count int64
		suite.repo.db.Model(&model.AttributeValue{}).Count(&count)
		suite.Equal(int64(2), count)
	})

	suite.Run("冲突更新已有记录", func() {
		updateData := []model.AttributeValue{
			{SkuID: 1001, KeyID: 1, Value: "Blue"},
		}

		err := suite.repo.SaveBatch(updateData)
		suite.NoError(err)

		// 验证数据更新
		var value model.AttributeValue
		suite.repo.db.Where("sku_id = ? AND key_id = ?", 1001, 1).First(&value)
		suite.Equal("Blue", value.Value)
	})
}

// --- GetIDsByKey 测试 ---
func (suite *AttributeValueRepoTestSuite) TestGetIDsByKey() {
	// 准备测试数据
	suite.createTestValues([]model.AttributeValue{
		{KeyID: 1, Value: "A"},
		{KeyID: 1, Value: "B"},
		{KeyID: 2, Value: "C"},
	})

	suite.Run("有效KeyID查询", func() {
		ids, err := suite.repo.GetIDsByKey(1)
		suite.NoError(err)
		suite.Len(ids, 2)
	})

	suite.Run("无效KeyID查询", func() {
		ids, err := suite.repo.GetIDsByKey(999)
		suite.NoError(err)
		suite.Empty(ids)
	})
}

// --- UpdateBatch 测试 ---
func (suite *AttributeValueRepoTestSuite) TestUpdateBatch() {
	// 准备初始数据
	initialData := []model.AttributeValue{
		{ID: 1, Value: "Old1"},
		{ID: 2, Value: "Old2"},
		{ID: 3, Value: "Old3"},
	}
	suite.createTestValues(initialData)

	// 更新数据
	updateData := []model.AttributeValue{
		{ID: 1, Value: "New1"},
		{ID: 3, Value: "New3"},
	}

	suite.Run("批量更新部分记录", func() {
		err := suite.repo.UpdateBatch(updateData)
		suite.NoError(err)

		// 验证更新结果
		var values []model.AttributeValue
		suite.repo.db.Find(&values)

		expected := map[int]string{
			1: "New1",
			2: "Old2",
			3: "New3",
		}
		for _, v := range values {
			suite.Equal(expected[v.ID], v.Value)
		}
	})
}

// --- DeleteBatch 测试 ---
func (suite *AttributeValueRepoTestSuite) TestDeleteBatch() {
	// 准备初始数据
	initialData := []model.AttributeValue{
		{ID: 1, Value: "A"},
		{ID: 2, Value: "B"},
		{ID: 3, Value: "C"},
	}
	suite.createTestValues(initialData)

	suite.Run("正常批量删除", func() {
		err := suite.repo.DeleteBatch([]int{1, 3})
		suite.NoError(err)

		// 验证剩余数据
		var remaining []model.AttributeValue
		suite.repo.db.Find(&remaining)
		suite.Len(remaining, 1)
		suite.Equal(2, remaining[0].ID)
	})

	suite.Run("删除不存在记录", func() {
		err := suite.repo.DeleteBatch([]int{999})
		suite.NoError(err)
	})
}

// 边界测试
func (suite *AttributeValueRepoTestSuite) TestEdgeCases() {
	suite.Run("空数据操作", func() {
		suite.Run("空批量保存", func() {
			err := suite.repo.SaveBatch(nil)
			suite.NoError(err)
		})

		suite.Run("空批量更新", func() {
			err := suite.repo.UpdateBatch(nil)
			suite.Error(err) // 应该返回错误
		})

		suite.Run("空批量删除", func() {
			err := suite.repo.DeleteBatch(nil)
			suite.NoError(err)
		})
	})

	suite.Run("大数据量测试", func() {
		// 生成1000条测试数据
		var bulkData []model.AttributeValue
		for i := 1; i <= 1000; i++ {
			bulkData = append(bulkData, model.AttributeValue{
				SkuID: uint64(i % 100),
				KeyID: uint64(i % 10),
				Value: fmt.Sprintf("Value%d", i),
			})
		}

		suite.Run("批量插入1000条", func() {
			err := suite.repo.SaveBatch(bulkData)
			suite.NoError(err)

			var count int64
			suite.repo.db.Model(&model.AttributeValue{}).Count(&count)
			suite.Equal(int64(1000), count)
		})
	})
}
