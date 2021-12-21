package system

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var ApplicationService = new(applicationService)

type applicationService struct{}

type ApplicationRegisterParams struct {
	Avatar          string
	Applicationname string
	Password        string
	NickName        string
}

//
// GetById
// @Description: 根据id获取用户信息
// @receiver applicationService
// @param Id
// @return data
// @return err
//
func (service *applicationService) GetById(Id uint) (data *system.Application, err error) {
	if Id == 0 {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	var entity system.Application
	var query = system.Application{Model: global.Model{Id: Id}}
	if err = global.Db.Model(query).First(&entity).Error; err != nil {
		return nil, gerror.NewCode(constants.CodeAppNotExists, "应用不存在")
	}
	return &entity, nil
}

type ApplicationFindParams struct {
	Id              uint   `json:"id" example:"7"`
	Applicationname string `json:"applicationname" example:"7"`
}

//
//  Find
//  @Description: 查询application
//  @receiver applicationService
//  @param applicationFindParams
//  @return data
//  @return err
//
func (service *applicationService) Find(applicationFindParams *ApplicationFindParams) (data *system.Application, err error) {
	var entity system.Application

	if applicationFindParams.Applicationname == "" && applicationFindParams.Id == 0 {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	var search = func(db *gorm.DB) *gorm.DB {
		if applicationFindParams.Id != 0 {
			db = db.Where("id = ?", applicationFindParams.Id)
		}
		if applicationFindParams.Applicationname != "" {
			db = db.Where("applicationname = ?", applicationFindParams.Applicationname)
		}
		return db
	}
	if err = global.Db.Scopes(search).Preload("Roles").First(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "用户查询失败!")
	}
	return &entity, nil
}

func (service *applicationService) GetByClientId(clientId string) (data *system.Application, err error) {
	if clientId == "" {
		return nil, gerror.NewCode(constants.CodeParamCantBlank)
	}
	var entity system.Application
	var query = system.Application{ClientId: clientId}
	if err = global.Db.Model(query).First(&entity).Error; err != nil {
		return nil, gerror.NewCode(constants.CodeAppNotExists, "应用不存在")
	}
	return &entity, nil
}
