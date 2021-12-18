package system

import (
	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/common"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var OperationRecord = new(operationRecord)

type operationRecord struct{}

type OperationRecordCreateParams struct {
	system.OperationRecord
}

// Create 创建记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) Create(info *OperationRecordCreateParams) error {
	return global.Db.Create(&info.OperationRecord).Error
}

// First 根据id获取单条操作记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) First(info *common.GetByID) (data *system.OperationRecord, err error) {
	var entity system.OperationRecord
	if err = global.Db.Where("id = ?", info.Id).First(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "查找记录失败")
	}
	return &entity, nil
}

// Delete 删除操作记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) Delete(info *common.GetByID) error {
	return global.Db.Delete(&system.OperationRecord{}, info.Id).Error
}

// Deletes 批量删除记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) Deletes(ids *common.GetByIDs) error {
	return global.Db.Delete(&[]system.OperationRecord{}, "id in (?)", ids.Ids).Error
}

type OperationRecordSearchParams struct {
	Path   string `json:"path" example:"请求路径"`
	Method string `json:"method" example:"请求方法"`
	Status int    `json:"status" example:"7"`
	common.PageInfo
}

// GetList 分页获取操作记录列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) GetList(info *OperationRecordSearchParams) (list []system.OperationRecord, total int64, err error) {
	db := global.Db.Model(&system.OperationRecord{})
	var entities []system.OperationRecord

	var search = func(db *gorm.DB) *gorm.DB { // 如果有条件搜索 下方会自动创建搜索语句
		if info.Method != "" {
			db = db.Where("method = ?", info.Method)
		}
		if info.Path != "" {
			db = db.Where("path LIKE ?", "%"+info.Path+"%")
		}
		if info.Status != 0 {
			db = db.Where("status = ?", info.Status)
		}
		return db.Order("id desc")
	}

	err = db.Scopes(search).Count(&total).Preload("User").Find(&entities).Error
	return entities, total, err
}
