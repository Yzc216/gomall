package repo

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"gorm.io/gorm"
	"log"
	"strings"
)

var (
	//判断结构体是否实现了Dark接口
	_ CategoryRepository = (*CategoryRepo)(nil) //把nil转化成*bird类型 赋值后即丢弃
)

type CategoryRepository interface {
	// 基础CRUD
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Category, error)
	ExistByName(ctx context.Context, name string) (bool, error)

	// 树形结构操作
	GetChildren(ctx context.Context, parentID uint64) ([]*model.Category, error)
	GetAncestors(ctx context.Context, id uint64) ([]*model.Category, error) // 获取所有祖先节点
	MoveBranch(ctx context.Context, id uint64, newParentID uint64) error    // 移动子树

	// SPU关联管理
	AddSPUs(ctx context.Context, categoryID uint64, spuIDs []uint64) error
	RemoveSPUs(ctx context.Context, categoryID uint64, spuIDs []uint64) error
	ClearSPUs(ctx context.Context, categoryID uint64) error
	GetSPUCount(ctx context.Context, categoryID uint64) (int64, error)

	// 批量操作
	BatchUpdateSort(ctx context.Context, ids []uint64, sorts map[uint64]int) error

	// 校验方法
	ValidateCategoryChain(ctx context.Context, ids []uint64) error // 验证分类层级连续性
	IsLeafCategory(ctx context.Context, id uint64) (bool, error)
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

// 基础CRUD操作
func (r *CategoryRepo) Create(ctx context.Context, c *model.Category) error {
	// 校验1：名称唯一性
	exist, err := r.ExistByName(ctx, c.Name)
	if err != nil {
		return fmt.Errorf("校验名称失败: %w", err)
	}
	if exist {
		return types.ErrDuplicateName
	}

	// 校验2：父分类有效性
	if c.ParentID != 0 {
		var parent model.Category
		if err := r.db.WithContext(ctx).First(&parent, c.ParentID).Error; err != nil {
			return fmt.Errorf("无效的父分类ID: %w", err)
		}
		c.Level = parent.Level + 1
	} else {
		c.Level = 1
	}

	return r.db.WithContext(ctx).Model(c).Create(c).Error
}

func (r *CategoryRepo) Update(ctx context.Context, c *model.Category) error {
	// 校验1：ID存在性
	var current model.Category
	if err := r.db.WithContext(ctx).First(&current, c.ID).Error; err != nil {
		return fmt.Errorf("分类不存在: %w", err)
	}

	// 校验2：名称修改时的唯一性
	if current.Name != c.Name {
		exist, err := r.ExistByName(ctx, c.Name)
		if err != nil {
			return err
		}
		if exist {
			return types.ErrDuplicateName
		}
	}

	// 校验3：父分类不能是自己
	if c.ParentID == c.ID {
		return types.ErrInvalidParent
	}

	return r.db.WithContext(ctx).Model(c).Save(c).Error
}

func (r *CategoryRepo) Delete(ctx context.Context, id uint64) error {
	// 校验1：存在性
	var c model.Category
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return fmt.Errorf("分类不存在: %w", err)
	}

	// 校验2：是否叶子节点
	if !c.IsLeaf {
		return types.ErrHasChildren
	}

	// 校验3：是否关联SPU
	count := r.db.Model(&c).Association("SPUs").Count()
	if count > 0 {
		return types.ErrHasAssociatedSPUs
	}

	return r.db.WithContext(ctx).Model(c).Delete(&c).Error
}

func (r *CategoryRepo) GetByID(ctx context.Context, id uint64) (*model.Category, error) {
	var c model.Category
	err := r.db.WithContext(ctx).First(&c, id).Error
	return &c, err
}

func (r *CategoryRepo) GetByName(ctx context.Context, name string) (*model.Category, error) {
	var c model.Category
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&c).Error
	return &c, err
}

func (r *CategoryRepo) ExistByID(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.Category{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *CategoryRepo) ExistByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.Category{}).
		Where("name = ?", name).
		Count(&count).Error
	return count > 0, err
}

// 树形结构操作（使用原生SQL优化）
const getChildrenSQL = `
WITH RECURSIVE cte AS (
    SELECT * FROM categories WHERE parent_id = ?
    UNION ALL
    SELECT c.* FROM categories c
    INNER JOIN cte ON c.parent_id = cte.id
)
SELECT * FROM cte ORDER BY level`

func (r *CategoryRepo) GetChildren(ctx context.Context, parentID uint64) ([]*model.Category, error) {
	var children []*model.Category
	err := r.db.WithContext(ctx).Raw(getChildrenSQL, parentID).Scan(&children).Error
	return children, err
}

const getAncestorsSQL = `
WITH RECURSIVE cte AS (
    SELECT * FROM categories WHERE id = ?
    UNION ALL
    SELECT c.* FROM categories c
    INNER JOIN cte ON c.id = cte.parent_id
)
SELECT * FROM cte WHERE id != ? ORDER BY level ASC`

func (r *CategoryRepo) GetAncestors(ctx context.Context, id uint64) ([]*model.Category, error) {
	var categories []*model.Category
	err := r.db.WithContext(ctx).Raw(getAncestorsSQL, id, id).Scan(&categories).Error
	return categories, err
}

const moveBranchSQL = `
WITH RECURSIVE cte AS (
    SELECT id, parent_id, level FROM categories WHERE id = ?
    UNION ALL
    SELECT c.id, c.parent_id, c.level FROM categories c
    INNER JOIN cte ON c.parent_id = cte.id
)
UPDATE categories SET 
    parent_id = CASE WHEN id = ? THEN ? ELSE parent_id END,
    level = CASE 
        WHEN id = ? THEN ?
        ELSE cte.level + ? 
    END
FROM cte WHERE categories.id = cte.id`

func (r *CategoryRepo) MoveBranch(ctx context.Context, id uint64, newParentID uint64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 获取新旧父节点信息
		var current, newParent model.Category
		if err := tx.First(&current, id).Error; err != nil {
			return err
		}
		if err := tx.First(&newParent, newParentID).Error; err != nil {
			return err
		}

		// 防止循环引用
		if isChild, err := r.isAncestor(tx, id, newParentID); err != nil || isChild {
			return fmt.Errorf("invalid parent category")
		}

		// 计算层级差
		levelDiff := newParent.Level + 1 - current.Level

		// 执行批量更新
		return tx.Exec(moveBranchSQL,
			id,
			id, newParentID,
			id, newParent.Level+1,
			levelDiff,
		).Error
	})
}

// SPU关联管理
func (r *CategoryRepo) AddSPUs(ctx context.Context, categoryID uint64, spuIDs []uint64) error {
	if len(spuIDs) == 0 {
		return nil
	}

	// 校验分类存在
	if exist, err := r.ExistByID(ctx, categoryID); err != nil || !exist {
		return types.ErrInvalidIDs
	}

	// 校验SPU存在
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.SPU{}).Where("id IN ?", spuIDs).Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(spuIDs) {
		return types.ErrSPUNotFound
	}

	// 批量插入（去重）
	values := make([]string, 0, len(spuIDs))
	for _, sid := range spuIDs {
		values = append(values, fmt.Sprintf("(%d, %d)", categoryID, sid))
	}

	sql := fmt.Sprintf(
		"INSERT IGNORE INTO spu_categories (category_id, spu_id) VALUES %s",
		strings.Join(values, ","),
	)

	return r.db.WithContext(ctx).Exec(sql).Error
}

func (r *CategoryRepo) RemoveSPUs(ctx context.Context, categoryID uint64, spuIDs []uint64) error {
	return r.db.WithContext(ctx).Exec(
		"DELETE FROM spu_categories WHERE category_id = ? AND spu_id IN ?",
		categoryID, spuIDs,
	).Error
}

func (r *CategoryRepo) ClearSPUs(ctx context.Context, categoryID uint64) error {
	return r.db.WithContext(ctx).Exec(
		"DELETE FROM spu_categories WHERE category_id = ?",
		categoryID,
	).Error
}

func (r *CategoryRepo) GetSPUCount(ctx context.Context, categoryID uint64) (int64, error) {
	var count int64
	// 直接统计中间表中关联该分类的记录数
	err := r.db.WithContext(ctx).Table("spu_categories").
		Where("category_id = ?", categoryID).
		Count(&count).
		Error
	return count, err
}

// 批量操作（事务处理）
func (r *CategoryRepo) BatchUpdateSort(ctx context.Context, ids []uint64, sorts map[uint64]int) error {
	if len(ids) == 0 {
		return nil
	}

	// 校验ID存在性
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Category{}).Where("id IN ?", ids).Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(ids) {
		return types.ErrInvalidIDs
	}

	// 构建CASE语句
	var caseStmt strings.Builder
	var params []interface{}

	caseStmt.WriteString("CASE id ")
	for _, id := range ids {
		caseStmt.WriteString("WHEN ? THEN ? ")
		params = append(params, id, sorts[id])
	}
	caseStmt.WriteString("END")

	// 生成 IN 占位符
	inClause := strings.Repeat("?,", len(ids)-1) + "?"

	// 将 ids 转换为 []interface{}
	idInterfaces := make([]interface{}, len(ids))
	for i, id := range ids {
		idInterfaces[i] = id
	}

	// 合并参数
	allParams := append(params, idInterfaces...)

	return r.db.WithContext(ctx).Exec(
		fmt.Sprintf("UPDATE categories SET sort = %s WHERE id IN ?",
			caseStmt.String(),
			inClause,
		),
		allParams...,
	).Error
}

// 校验方法
func (r *CategoryRepo) ValidateCategoryChain(ctx context.Context, ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}

	// 获取所有分类的层级信息
	var categories []*model.Category
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return err
	}

	// 构建层级映射
	levelMap := make(map[uint64]int8)
	for _, c := range categories {
		levelMap[c.ID] = c.Level
	}

	// 验证层级连续性
	for i := 1; i < len(ids); i++ {
		prevLevel := levelMap[ids[i-1]]
		currLevel := levelMap[ids[i]]
		if currLevel != prevLevel+1 {
			return types.ErrInvalidCategoryChain
		}
	}

	return nil
}

func (r *CategoryRepo) IsLeafCategory(ctx context.Context, id uint64) (bool, error) {
	// 1. 校验分类存在性
	if exist, err := r.ExistByID(ctx, id); err != nil || !exist {
		return false, types.ErrInvalidIDs
	}

	// 2. 检查子分类是否存在（利用索引优化）
	var childCount int64
	err := r.db.WithContext(ctx).
		Model(&model.Category{}).
		Where("parent_id = ?", id).
		Count(&childCount).Error
	if err != nil {
		return false, fmt.Errorf("查询子分类失败: %w", err)
	}

	// 3. 双重验证（可选）
	// 如果结构体有IsLeaf字段，可以与数据库查询结果比对
	var c model.Category
	if err := r.db.WithContext(ctx).Select("is_leaf").First(&c, id).Error; err == nil {
		if c.IsLeaf != (childCount == 0) {
			log.Printf("数据不一致告警: 分类%d的IsLeaf字段异常，实际子节点数=%d", id, childCount)
		}
	}

	return childCount == 0, nil
}

// 辅助方法
func (r *CategoryRepo) isAncestor(tx *gorm.DB, id uint64, ancestorID uint64) (bool, error) {
	var count int64
	err := tx.Raw(
		"WITH RECURSIVE cte AS ("+
			"SELECT id, parent_id FROM categories WHERE id = ? "+
			"UNION ALL "+
			"SELECT c.id, c.parent_id FROM categories c "+
			"INNER JOIN cte ON c.parent_id = cte.id"+
			") SELECT COUNT(*) FROM cte WHERE id = ?",
		id, ancestorID,
	).Scan(&count).Error

	return count > 0, err
}

//func (s *ProductService) CreateSPUWithCategories(spu *model.SPU, categoryIDs []uint64) error {
//	// 步骤1：验证分类有效性
//	if err := s.validateCategories(categoryIDs); err != nil {
//		return err
//	}
//
//	// 步骤2：开启事务
//	return s.txExecutor.Transaction(func(tx *gorm.DB) error {
//		// 步骤3：创建SPU
//		if err := tx.Create(spu).Error; err != nil {
//			return err
//		}
//
//		// 步骤4：建立分类关联
//		return s.CategoryRepo.AddSPUs(spu.ID, categoryIDs)
//	})
//}
//
//// 分类验证逻辑
//func (s *ProductService) validateCategories(ids []uint64) error {
//	for _, id := range ids {
//		// 检查是否为叶子分类
//		if isLeaf, err := s.CategoryRepo.IsLeafCategory(id); err != nil || !isLeaf {
//			return fmt.Errorf("invalid category %d", id)
//		}
//
//		// 检查分类状态有效性
//		if _, err := s.CategoryRepo.GetByID(id); err != nil {
//			return err
//		}
//	}
//	return nil
//}
