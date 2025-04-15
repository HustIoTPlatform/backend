package service

import (
	"time"

	"github.com/Thingsly/backend/internal/dal"
	model "github.com/Thingsly/backend/internal/model"

	"github.com/spf13/viper"
)

type AutomateTask struct {
}

// OnceTaskExecute
// @description
// @params t *AutomateTask
// @return error
func (t *AutomateTask) OnceTaskExecute() error {
	limit := viper.GetInt("automation_task_confg.once_task_limit")
	result, err := dal.GetOnceTaskListWithLock(limit)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	var expIds []string
	for _, v := range result {
		if v.ExpirationTime > 0 && v.ExecutionTime.Add(time.Duration(v.ExpirationTime)*time.Minute).Before(time.Now()) {
			expIds = append(expIds, v.ID)
			continue
		}
		t.TaskAutomationExecute(v.SceneAutomationID)
	}
	return dal.TaskExpirationSave(expIds)
}

// TaskAutomationExecute
// @description
// @params t *AutomateTask
// @return error
func (*AutomateTask) TaskAutomationExecute(sceneAutomationId string) {

	if GroupApp.CheckSceneAutomationHasClose(sceneAutomationId) {
		return
	}
	actions, err := dal.GetActionInfoListBySceneAutomationId([]string{sceneAutomationId})
	if err != nil {
		return
	}
	var (
		deviceIds      []string
		deviceConfigId []string
	)
	for _, v := range actions {
		if v.ActionType == model.AUTOMATE_ACTION_TYPE_MULTIPLE && v.ActionTarget != nil {
			deviceConfigId = append(deviceConfigId, *v.ActionTarget)
		}
	}
	if len(deviceConfigId) > 0 {
		deviceIds, err = dal.GetDeviceIdsByDeviceConfigId(deviceConfigId)
		if err != nil {
			return
		}
	}
	_ = GroupApp.SceneAutomateExecute(sceneAutomationId, deviceIds, actions)

}

// PeriodicTaskExecute
// @description
// @params t *AutomateTask
// @return error
func (t *AutomateTask) PeriodicTaskExecute() error {

	limit := viper.GetInt("automation_task_confg.periodic_task_limit")
	result, err := dal.GetPeriodicTaskListWithLock(limit)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return nil
	}
	for _, v := range result {
		if v.ExpirationTime > 0 && v.ExecutionTime.Add(time.Duration(v.ExpirationTime)*time.Minute).Before(time.Now()) {
			continue
		}
		t.TaskAutomationExecute(v.SceneAutomationID)
	}
	return nil
}
