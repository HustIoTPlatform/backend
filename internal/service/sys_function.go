package service

import (
	"github.com/Thingsly/backend/internal/dal"
	model "github.com/Thingsly/backend/internal/model"
	"github.com/Thingsly/backend/pkg/errcode"
	"github.com/Thingsly/backend/pkg/global"
)

type SysFunction struct{}

func (*SysFunction) GetSysFuncion(lang string) ([]*model.SysFunction, error) {
	data, err := dal.GetAllSysFunction()

	for _, v := range data {
		description := global.ResponseHandler.ErrManager.GetMessageStr(*v.Description, lang)
		v.Description = &description
	}
	return data, err
}

func (*SysFunction) UpdateSysFuncion(function_id string) error {
	old, err := dal.GetSysFunctionById(function_id)
	if err != nil {
		return errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	if old.ID == "" {
		return errcode.WithData(errcode.CodeSystemError, map[string]interface{}{
			"msg": "id is nil",
		})
	}

	var upTarget string

	if old.EnableFlag == "enable" {
		upTarget = "disable"
	} else {
		upTarget = "enable"
	}

	err = dal.UpdateSysFunction(function_id, upTarget)
	if err != nil {
		return errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return err
}
