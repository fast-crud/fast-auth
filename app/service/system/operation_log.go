package system

import (
	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/common"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var OperationLog = new(operationLog)

type operationLog struct{}

type OperationLogCreateParams struct {
	system.OperationLog
}

// Create 创建记录

func (s *operationLog) Create(info *OperationLogCreateParams) error {
	return global.Db.Create(&info.OperationLog).Error
}

// First 根据id获取单条操作记录

func (s *operationLog) First(info *common.GetByID) (data *system.OperationLog, err error) {
	var entity system.OperationLog
	if err = global.Db.Where("id = ?", info.Id).First(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "查找记录失败")
	}
	return &entity, nil
}

// Delete 删除操作记录

func (s *operationLog) Delete(info *common.GetByID) error {
	return global.Db.Delete(&system.OperationLog{}, info.Id).Error
}

// Deletes 批量删除记录

func (s *operationLog) Deletes(ids *common.GetByIDs) error {
	return global.Db.Delete(&[]system.OperationLog{}, "id in (?)", ids.Ids).Error
}

type OperationRecordSearchParams struct {
	Path   string `json:"path" example:"请求路径"`
	Method string `json:"method" example:"请求方法"`
	Status int    `json:"status" example:"7"`
	common.PageInfo
}

// GetList 分页获取操作记录列表

func (s *operationLog) GetList(info *OperationRecordSearchParams) (list []system.OperationLog, total int64, err error) {
	db := global.Db.Model(&system.OperationLog{})
	var entities []system.OperationLog

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
