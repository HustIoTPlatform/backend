package subscribe

import (
	"encoding/json"
	"strings"
	"time"

	initialize "github.com/HustIoTPlatform/backend/initialize"
	dal "github.com/HustIoTPlatform/backend/internal/dal"
	"github.com/HustIoTPlatform/backend/internal/model"
	service "github.com/HustIoTPlatform/backend/internal/service"

	"github.com/go-basic/uuid"
	"github.com/sirupsen/logrus"
)

func DeviceEvent(payload []byte, topic string) (string, string, string, error) {
	var messageId string
	topicList := strings.Split(topic, "/")
	if len(topicList) < 3 {
		messageId = ""
	} else {
		messageId = topicList[2]
	}

	eventPayload, err := verifyPayload(payload)
	if err != nil {
		logrus.Error(err.Error())
		return "", "", "", err
	}

	device, err := initialize.GetDeviceCacheById(eventPayload.DeviceId)
	if err != nil {
		logrus.Error(err.Error())
		return "", "", "", err
	}

	logrus.Debug("event message:", eventPayload)

	eventValues, err := verifyEventPayload(eventPayload.Values)
	if err != nil {
		logrus.Error(err.Error())
		return device.DeviceNumber, messageId, "", err
	}
	logrus.Debug("event message:", eventValues)

	err = deviceEventHandle(device, eventValues, topic)
	if err != nil {
		logrus.Error(err.Error())
		return device.DeviceNumber, messageId, "", err
	}
	return device.DeviceNumber, messageId, eventValues.Method, nil

}

func deviceEventHandle(device *model.Device, eventValues *model.EventInfo, topic string) error {

	if device.DeviceConfigID != nil && *device.DeviceConfigID != "" {
		eventValuesByte, err := json.Marshal(eventValues)
		if err != nil {
			logrus.Error("JSON marshaling failed:", err)
			return err
		}
		neweventValues, err := service.GroupApp.DataScript.Exec(device, "F", eventValuesByte, topic)
		if err != nil {
			logrus.Error("Error in event script processing: ", err.Error())
		}
		if neweventValues != nil {
			err = json.Unmarshal(neweventValues, &eventValues)
			if err != nil {
				logrus.Error("Error in attribute script processing: ", err.Error())
			}
		}
	}

	paramsJsonBytes, err := json.Marshal(eventValues.Params)
	if err != nil {
		logrus.Fatalf("JSON marshaling failed: %s", err)
		return err
	}
	paramsJsonString := string(paramsJsonBytes)
	eventDatas := &model.EventData{
		ID:       uuid.New(),
		DeviceID: device.ID,
		Identify: eventValues.Method,
		T:        time.Now().UTC(),
		Datum:    &paramsJsonString,
		TenantID: &device.TenantID,
	}

	go func() {

		err = service.GroupApp.Execute(device, service.AutomateFromExt{
			TriggerParamType: model.TRIGGER_PARAM_TYPE_EVT,
			TriggerParam:     []string{eventValues.Method},
			TriggerValues: map[string]interface{}{
				eventValues.Method: paramsJsonString,
			},
		})
		if err != nil {
			logrus.Error("Automation execution failed, err:", err)
		}
	}()
	err = dal.CreateEventData(eventDatas)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return err
}
