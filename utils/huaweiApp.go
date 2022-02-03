package utils

import (
	"context"
	"encoding/json"

	pb "github.com/nostressdev/notifications/proto"
	pushkit "github.com/nostressdev/notifications/pushkit/push/core"
	"github.com/nostressdev/notifications/pushkit/push/model"
)

type HuaweiApp struct {
	App *pushkit.HttpPushClient
}

func (app *HuaweiApp) SendMessage(request *pb.SendDevicePushRequest, device *pb.Device) (string, error) {
	lang := device.DeviceInfo.Language
	systemName := getSystemName(device.DeviceInfo.DeviceType)
	data, ok := request.Notification.Data[systemName]
	if !ok {
		data, ok = request.Notification.Data["common"]
		if !ok {
			data = &pb.JSONObject{Value: make(map[string]string)}
		}
	}
	data.Value["on_click"] = request.Notification.ClickAction[systemName]
	if data.Value["on_click"] == "" {
		data.Value["on_click"] = request.Notification.ClickAction["common"]
	}
	bytes, err := json.Marshal(data.Value)
	if err != nil {
		return "", err
	}
	msg := &model.MessageRequest{
		Message: &model.Message{
			Notification: &model.Notification{
				Title: localize(request.Notification.Title, lang),
				Body:  localize(request.Notification.Body, lang),
				Image: request.Notification.ImageURL,
			},
			Android: &model.AndroidConfig{},
			Token:   []string{device.DeviceInfo.Identifier},
			Data:    string(bytes),
		},
	}
	response, err := app.App.SendMessage(context.Background(), msg)
	if err != nil {
		return "", err
	}
	return response.Msg, nil
}
