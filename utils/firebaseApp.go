package utils

import (
	"context"
	"log"

	pb "github.com/nostressdev/notifications/proto"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

type FirebaseApp struct {
	App *firebase.App
}

func localize(dict map[string]string, lang string) string {
	if lang == "" || dict[lang] == "" {
		return dict["en"]
	}
	return dict[lang]
}

func getSystemName(deviceType pb.DeviceType) string {
	if deviceType == pb.DeviceType_ANDROID {
		return "android"
	} else if deviceType == pb.DeviceType_HUAWEI {
		return "huawei"
	} else if deviceType == pb.DeviceType_IOS {
		return "apns"
	} else if deviceType == pb.DeviceType_WEB {
		return "webpush"
	} else if deviceType == pb.DeviceType_EMAIL {
		return "email"
	} else {
		return ""
	}
}

func (app *FirebaseApp) SendMessage(request *pb.SendDevicePushRequest, device *pb.Device) (string, error) {
	lang := device.DeviceInfo.Language
	client, err := app.App.Messaging(context.Background())
	if err != nil {
		return "", err
	}

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
	log.Println(device.DeviceInfo.Identifier)
	log.Println(&messaging.Notification{
		Title:    localize(request.Notification.Title, lang),
		Body:     localize(request.Notification.Body, lang),
		ImageURL: request.Notification.ImageURL,
	})
	
	response, err := client.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title:    localize(request.Notification.Title, lang),
			Body:     localize(request.Notification.Body, lang),
			ImageURL: request.Notification.ImageURL,
		},
		Android: &messaging.AndroidConfig{},
		Token:   device.DeviceInfo.Identifier,
		Data:    data.Value,
	})
	if err != nil {
		return "", err
	}
	return response, nil
}
