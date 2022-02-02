package notifications

import (
	"log"
	pb "github.com/nostressdev/notifications/proto"
)

type Service interface {
	CreateVirtualUser(accountID string) (string, error)
	GetUser(string) (*pb.User, error)
	AddUserTag(string, string) error
	DeleteUserTag(string, string) error
	GetUsersByTag(string) ([]*pb.User, error)
	AddDevice(*pb.DeviceInfo, string) (string, error)
	GetDevice(string) (*pb.Device, error)
	DeleteDevice(string) error
}

func NewFDB(config *ConfigFDB) Service {
	if config == nil {
		log.Fatalln("fdb config must not be nil")
	}
	return &implFDB{
		ConfigFDB: config,
		UsersSubspace: config.Subspace.Sub("users"),
		DevicesSubspace: config.Subspace.Sub("devices"),
		UserTagsSubspace: config.Subspace.Sub("user_tags"),
		TagUsersSubspace: config.Subspace.Sub("tag_users"),
	}
}
