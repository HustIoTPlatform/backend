package service

import (
	dal "github.com/Thingsly/backend/internal/dal"
	model "github.com/Thingsly/backend/internal/model"
	"github.com/Thingsly/backend/pkg/errcode"

	"github.com/sirupsen/logrus"
)

type Logo struct{}

func (*Logo) UpdateLogo(UpdateLogoReq *model.UpdateLogoReq) error {
	condsMap, err := StructToMapAndVerifyJson(UpdateLogoReq)
	if err != nil {
		return errcode.WithData(errcode.CodeParamError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	err = dal.UpdateLogo(UpdateLogoReq.Id, condsMap)
	if err != nil {
		logrus.Error(err)
		return errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"err": err.Error(),
		})
	}
	return err
}

func (*Logo) GetLogoList() (map[string]interface{}, error) {

	total, list, err := dal.GetLogoList()
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"err": err.Error(),
		})
	}
	logoListRsp := make(map[string]interface{})
	logoListRsp["total"] = total
	logoListRsp["list"] = list

	return logoListRsp, err
}
