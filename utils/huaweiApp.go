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
	data, ok := request.Notification.Data["huawei"]
	if !ok {
		data = request.Notification.Data["common"]
	}
	data.Value["on_click"] = request.Notification.ClickAction["huawei"]
	if data.Value["on_click"] == "" {
		data.Value["on_click"] = request.Notification.ClickAction["common"]
	}
	payload, err := json.Marshal(data.Value)
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
			Data:    string(payload),
		},
	}
	response, err := app.App.SendMessage(context.Background(), msg)
	if err != nil {
		return "", err
	}
	return response.Msg, nil
}
