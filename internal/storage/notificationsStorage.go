package storage

import (
	pb "github.com/nostressdev/notifications/proto"
)

type NotificationsStorage interface {
	CreateVirtualUser(accountID string) (string, error)
	GetUser(string) (*pb.User, error)
	AddUserTag(string, string) error
	DeleteUserTag(string, string) error
	GetUsersByTag(string) ([]*pb.User, error)
	AddDevice(info *pb.DeviceInfo, userID, deviceID string) (string, error)
	GetDevice(userID, deviceID string) (*pb.Device, error)
	DeleteDevice(userID, deviceID string) error
}
