package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Thingsly/backend/internal/dal"
	model "github.com/Thingsly/backend/internal/model"
	"github.com/Thingsly/backend/pkg/errcode"

	"github.com/go-basic/uuid"
	"github.com/sirupsen/logrus"
)

type Alarm struct{}

// CreateAlarmConfig
func (*Alarm) CreateAlarmConfig(req *model.CreateAlarmConfigReq) (data *model.AlarmConfig, err error) {
	data = &model.AlarmConfig{}
	t := time.Now().UTC()
	data.ID = uuid.New()
	data.Name = req.Name
	data.Description = req.Description
	data.AlarmLevel = req.AlarmLevel
	data.NotificationGroupID = req.NotificationGroupID
	data.CreatedAt = t
	data.UpdatedAt = t
	data.TenantID = req.TenantID
	data.Remark = req.Remark
	data.Enabled = req.Enabled

	err = dal.CreateAlarmConfig(data)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return
}

// DeleteAlarmConfig
func (*Alarm) DeleteAlarmConfig(id string) (err error) {
	err = dal.DeleteAlarmConfig(id)
	if err != nil {
		return errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return
}

// UpdateAlarmConfig
func (*Alarm) UpdateAlarmConfig(req *model.UpdateAlarmConfigReq) (data *model.AlarmConfig, err error) {
	data = &model.AlarmConfig{}
	data.ID = req.ID
	if req.Name != nil {
		data.Name = *req.Name
	}
	if req.Description != nil {
		data.Description = req.Description
	}
	if req.AlarmLevel != nil {
		data.AlarmLevel = *req.AlarmLevel
	}
	if req.NotificationGroupID != nil {
		data.NotificationGroupID = *req.NotificationGroupID
	}
	data.UpdatedAt = time.Now().UTC()
	data.TenantID = *req.TenantID
	data.Remark = req.Remark
	if req.Enabled != nil {
		data.Enabled = *req.Enabled
	}

	err = dal.UpdateAlarmConfig(data)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	data, err = dal.GetAlarmByID(req.ID)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return data, nil
}

// GetAlarmConfigListByPage
func (*Alarm) GetAlarmConfigListByPage(req *model.GetAlarmConfigListByPageReq) (data map[string]interface{}, err error) {

	total, list, err := dal.GetAlarmConfigListByPage(req)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	data = make(map[string]interface{})
	data["total"] = total
	data["list"] = list
	return
}

// UpdateAlarmInfo
func (*Alarm) UpdateAlarmInfo(req *model.UpdateAlarmInfoReq, userid string) (alarmInfo *model.AlarmInfo, err error) {

	alarmInfo, err = dal.GetAlarmInfoByID(req.Id)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	alarmInfo.Processor = &userid
	if req.ProcessingResult != nil && *req.ProcessingResult != "" {
		alarmInfo.ProcessingResult = *req.ProcessingResult
	}
	err = dal.UpdateAlarmInfo(alarmInfo)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return
}

// UpdateAlarmInfoBatch
func (*Alarm) UpdateAlarmInfoBatch(req *model.UpdateAlarmInfoBatchReq, userid string) error {
	if len(req.Id) == 0 {
		return errcode.WithData(errcode.CodeParamError, map[string]interface{}{
			"id": "id is empty",
		})
	}
	err := dal.UpdateAlarmInfoBatch(req, userid)
	if err != nil {
		return errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return err
}

// GetAlarmInfoListByPage
func (*Alarm) GetAlarmInfoListByPage(req *model.GetAlarmInfoListByPageReq) (data map[string]interface{}, err error) {

	total, list, err := dal.GetAlarmInfoListByPage(req)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	data = make(map[string]interface{})
	data["total"] = total
	data["list"] = list
	return
}

// GetAlarmHisttoryListByPage
func (*Alarm) GetAlarmHisttoryListByPage(req *model.GetAlarmHisttoryListByPage, tenantID string) (data map[string]interface{}, err error) {

	total, list, err := dal.GetAlarmHistoryListByPage(req, tenantID)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	data = make(map[string]interface{})
	data["total"] = total
	data["list"] = list
	return
}
func (*Alarm) AlarmHistoryDescUpdate(req *model.AlarmHistoryDescUpdateReq, tenantID string) (err error) {
	err = dal.AlarmHistoryDescUpdate(req, tenantID)
	if err != nil {
		return errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return
}
func (*Alarm) GetDeviceAlarmStatus(req *model.GetDeviceAlarmStatusReq) bool {
	return dal.GetDeviceAlarmStatus(req)
}

func (*Alarm) GetConfigByDevice(req *model.GetDeviceAlarmStatusReq) ([]model.AlarmConfig, error) {
	data, err := dal.GetConfigByDevice(req)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return data, nil
}

// AddAlarmInfo
func (*Alarm) AddAlarmInfo(alarmConfigID, content string) (bool, string) {

	alarmConfig, err := dal.GetAlarmByID(alarmConfigID)
	if err != nil {
		logrus.Error(err)
		return false, ""
	}

	if alarmConfig.Enabled != "Y" {
		return false, ""
	}

	if alarmConfig.NotificationGroupID != "" {
		title := alarmConfig.Name + "[" + alarmConfig.AlarmLevel + "]" + time.Now().Format("2006-01-02 15:04:05")
		GroupApp.NotificationServicesConfig.ExecuteNotification(alarmConfig.NotificationGroupID, title, content)
	}

	id := uuid.New()
	t := time.Now().UTC()
	err = dal.CreateAlarmInfo(&model.AlarmInfo{
		ID:               id,
		Name:             alarmConfig.Name,
		AlarmConfigID:    alarmConfigID,
		AlarmLevel:       &alarmConfig.AlarmLevel,
		Content:          &content,
		AlarmTime:        t,
		Description:      alarmConfig.Description,
		ProcessingResult: "UND",
		TenantID:         alarmConfig.TenantID,
	})
	if err != nil {
		logrus.Error(err)
		return false, ""
	}
	return true, id
}

func (*Alarm) AlarmRecovery(alarmConfigID, content, scene_automation_id, group_id string, device_ids []string) (bool, string) {
	alarmConfig, err := dal.GetAlarmByID(alarmConfigID)
	if err != nil {
		logrus.Error(err)
		return false, ""
	}

	device_ids_str, _ := json.Marshal(device_ids)
	id := uuid.New()
	t := time.Now().UTC()
	err = dal.AlarmHistorySave(&model.AlarmHistory{
		ID:                id,
		Name:              alarmConfig.Name,
		AlarmConfigID:     alarmConfigID,
		Content:           &content,
		Description:       alarmConfig.Description,
		TenantID:          alarmConfig.TenantID,
		SceneAutomationID: scene_automation_id,
		GroupID:           group_id,
		AlarmDeviceList:   string(device_ids_str),
		AlarmStatus:       "N",
		CreateAt:          t,
	})
	if err != nil {
		logrus.Error(err)
		return false, ""
	}
	return true, id
}

func (*Alarm) AlarmExecute(alarmConfigID, content, scene_automation_id, group_id string, device_ids []string) (bool, string, string) {
	var alarmName string
	alarmConfig, err := dal.GetAlarmByID(alarmConfigID)
	if err != nil {
		logrus.Error(err)
		return false, alarmName, ""
	}

	if alarmConfig.Enabled != "Y" {
		return false, alarmName, "alarm is not enabled"
	}
	alarmName = alarmConfig.Name
	if alarmConfig.NotificationGroupID != "" {
		title := alarmConfig.Name + "[" + alarmConfig.AlarmLevel + "]" + time.Now().Format("2006-01-02 15:04:05")
		GroupApp.NotificationServicesConfig.ExecuteNotification(alarmConfig.NotificationGroupID, title, content)
	}
	device_ids_str, _ := json.Marshal(device_ids)
	id := uuid.New()
	t := time.Now().UTC()
	err = dal.AlarmHistorySave(&model.AlarmHistory{
		ID:                id,
		Name:              alarmConfig.Name,
		AlarmConfigID:     alarmConfigID,
		Content:           &content,
		Description:       alarmConfig.Description,
		TenantID:          alarmConfig.TenantID,
		SceneAutomationID: scene_automation_id,
		GroupID:           group_id,
		AlarmDeviceList:   string(device_ids_str),
		AlarmStatus:       alarmConfig.AlarmLevel,
		CreateAt:          t,
	})
	if err != nil {
		logrus.Error(err)
		return false, alarmName, ""
	}
	for _, deviceId := range device_ids {
		deviceInfo, _ := dal.GetDeviceByID(deviceId)
		go GroupApp.AlarmMessagePushSend(alarmConfig.Name, id, deviceInfo)
	}
	return true, alarmName, ""
}

func (*Alarm) GetAlarmInfoHistoryByID(id string) (map[string]interface{}, error) {
	alarmInfo, err := dal.GetAlarmInfoHistoryByID(id)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"sql_error": err.Error(),
		})
	}
	return alarmInfo, nil
}

func (a *Alarm) GetAlarmDeviceCountsByTenant(tenantID string) (*model.AlarmDeviceCountsResponse, error) {
	ctx := context.Background()
	db := &dal.LatestDeviceAlarmQuery{}

	totalCount, err := db.CountDevicesByTenantAndStatus(ctx, tenantID)
	if err != nil {
		return nil, errcode.WithData(errcode.CodeDBError, map[string]interface{}{
			"operation": "count_alarm_devices",
			"error":     err.Error(),
		})
	}

	return &model.AlarmDeviceCountsResponse{
		AlarmDeviceTotal: int64(totalCount),
	}, nil
}
