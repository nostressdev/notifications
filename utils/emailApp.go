package utils

import (
	pb "github.com/nostressdev/notifications/proto"

	gomail "gopkg.in/gomail.v2"
)

type SMTP struct {
	Host     string
	Port     int
	Login    string
	Password string
	Sender   string
}

type EmailApp struct {
	SMTP *SMTP
}

func (app *EmailApp) SendMessage(request *pb.SendDevicePushRequest, device *pb.Device) (string, error) {
	lang := device.DeviceInfo.Language
	msg := gomail.NewMessage()
	msg.SetHeader("From", app.SMTP.Sender)
	msg.SetHeader("To", device.DeviceInfo.Identifier)
	msg.SetHeader("Subject", localize(request.Notification.Title, lang))
	msg.SetBody("text/html", localize(request.Notification.Body, lang))
	dialer := gomail.NewDialer(app.SMTP.Host, app.SMTP.Port, app.SMTP.Login, app.SMTP.Password)
	if err := dialer.DialAndSend(msg); err != nil {
		return "", err
	}
	return "", nil
}
