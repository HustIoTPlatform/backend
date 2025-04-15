package api

import (
	"github.com/Thingsly/backend/internal/model"
	"github.com/Thingsly/backend/internal/service"
	common "github.com/Thingsly/backend/pkg/common"
	"github.com/Thingsly/backend/pkg/errcode"
	"github.com/Thingsly/backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SceneAutomationsApi struct{}

// Create scene automation
// /api/v1/scene_automations [post]
func (*SceneAutomationsApi) CreateSceneAutomations(c *gin.Context) {
	logrus.Info("Create scene automation")
	var req model.CreateSceneAutomationReq
	if !BindAndValidate(c, &req) {
		return
	}
	userClaims := c.MustGet("claims").(*utils.UserClaims)

	id, err := service.GroupApp.SceneAutomation.CreateSceneAutomation(&req, userClaims)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", map[string]interface{}{"scene_automation_id": id})
}

// Delete scene automation
// /api/v1/scene_automations/{id} [delete]
func (*SceneAutomationsApi) DeleteSceneAutomations(c *gin.Context) {
	id := c.Param("id")
	err := service.GroupApp.SceneAutomation.DeleteSceneAutomation(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", 1)
}

// Update scene automation
// /api/v1/scene_automations/switch/{id} [post]
func (*SceneAutomationsApi) SwitchSceneAutomations(c *gin.Context) {
	id := c.Param("id")
	err := service.GroupApp.SceneAutomation.SwitchSceneAutomation(id, "")
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", nil)
}

// Update scene automation
// /api/v1/scene_automation [put]
func (*SceneAutomationsApi) UpdateSceneAutomations(c *gin.Context) {
	var req model.UpdateSceneAutomationReq
	if !BindAndValidate(c, &req) {
		return
	}
	userClaims := c.MustGet("claims").(*utils.UserClaims)
	id, err := service.GroupApp.SceneAutomation.UpdateSceneAutomation(&req, userClaims)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", map[string]interface{}{"scene_automation_id": id})
}

// Get scene automation by id
// /api/v1/scene_automations/detail/{id} [get]
func (*SceneAutomationsApi) HandleSceneAutomations(c *gin.Context) {
	id := c.Param("id")
	data, err := service.GroupApp.SceneAutomation.GetSceneAutomation(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", data)
}

// Get scene automation by page
// /api/v1/scene_automations/list [get]
func (*SceneAutomationsApi) HandleSceneAutomationsByPage(c *gin.Context) {
	var req model.GetSceneAutomationByPageReq
	if !BindAndValidate(c, &req) {
		return
	}
	userClaims := c.MustGet("claims").(*utils.UserClaims)
	data, err := service.GroupApp.SceneAutomation.GetSceneAutomationByPageReq(&req, userClaims)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", data)
}

// Get scene automation with alarm by page
// /api/v1/scene_automations/alarm [get]
func (*SceneAutomationsApi) HandleSceneAutomationsWithAlarmByPage(c *gin.Context) {
	var req model.GetSceneAutomationsWithAlarmByPageReq
	if !BindAndValidate(c, &req) {
		return
	}

	if common.IsStringEmpty(req.DeviceId) && common.IsStringEmpty(req.DeviceConfigId) {
		c.Error(errcode.WithData(errcode.CodeParamError, map[string]interface{}{
			"error": "device_id and device_config_id can not be empty at the same time",
		}))
		return
	}
	userClaims := c.MustGet("claims").(*utils.UserClaims)
	data, err := service.GroupApp.SceneAutomation.GetSceneAutomationWithAlarmByPageReq(&req, userClaims)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", data)
}

// Get scene automation log by page
// /api/v1/scene_automations/log [get]
func (*SceneAutomationsApi) HandleSceneAutomationsLog(c *gin.Context) {
	var req model.GetSceneAutomationLogReq
	if !BindAndValidate(c, &req) {
		return
	}
	userClaims := c.MustGet("claims").(*utils.UserClaims)
	data, err := service.GroupApp.SceneAutomationLog.GetSceneAutomationLog(&req, userClaims)
	if err != nil {
		c.Error(err)
		return
	}
	c.Set("data", data)
}
