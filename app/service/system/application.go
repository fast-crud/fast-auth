package system

import (
	"github.com/fast-crud/fast-auth/app/constants"
	"github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/gogf/gf/v2/errors/gerror"
)

var ApplicationService = new(applicationService)

type applicationService struct{}

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
