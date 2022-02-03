package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	pb "github.com/nostressdev/notifications/proto"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type ConfigNotificationsFDB struct {
	DB       fdb.Database
	Subspace subspace.Subspace
}

type NotificationsStorageFDB struct {
	*ConfigNotificationsFDB
	UsersSubspace    subspace.Subspace
	DevicesSubspace  subspace.Subspace
	UserTagsSubspace subspace.Subspace
	TagUsersSubspace subspace.Subspace
}


func NewNotificationsFDB(config *ConfigNotificationsFDB) NotificationsStorage {
	if config == nil {
		log.Fatalln("fdb config must not be nil")
	}
	return &NotificationsStorageFDB{
		ConfigNotificationsFDB: config,
		UsersSubspace: config.Subspace.Sub("users"),
		DevicesSubspace: config.Subspace.Sub("devices"),
		UserTagsSubspace: config.Subspace.Sub("user_tags"),
		TagUsersSubspace: config.Subspace.Sub("tag_users"),
	}
}


func (s *NotificationsStorageFDB) CreateVirtualUser(accountID string) (string, error) {
	user := &pb.User{
		AccountID: accountID,
	}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		userBytes, err := proto.Marshal(user)
		if err != nil {
			return nil, err
		}
		tr.Set(s.UsersSubspace.Sub(user.AccountID), userBytes)
		return nil, nil
	})
	return user.AccountID, err
}

func (s *NotificationsStorageFDB) GetUser(AccountID string) (*pb.User, error) {
	user := new(pb.User)
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		userBytes, err := tr.Get(s.UsersSubspace.Sub(AccountID)).Get()
		if err != nil {
			return nil, err
		}
		err = proto.Unmarshal(userBytes, user)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return user, err
}

func (s *NotificationsStorageFDB) AddUserTag(userID, tag string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		user, err := s.GetUser(userID)
		if err != nil {
			return nil, err
		} else if user.AccountID != userID {
			return nil, errors.New(fmt.Sprintf("no such user %s", userID))
		}

		userTagsBytes, err := tr.Get(s.UserTagsSubspace.Sub(userID)).Get()
		userTags := []string{}
		if len(userTagsBytes) > 0 {
			if err := json.Unmarshal(userTagsBytes, &userTags); err != nil {
				return nil, err
			}
		}
		for _, userTag := range userTags {
			if userTag == tag {
				return nil, nil
			}
		}
		userTags = append(userTags, tag)
		userTagsBytes, err = json.Marshal(userTags)
		if err != nil {
			return nil, err
		}
		tr.Set(s.UserTagsSubspace.Sub(userID), userTagsBytes)

		tagUsersBytes, err := tr.Get(s.TagUsersSubspace.Sub(tag)).Get()
		tagUsers := []string{}
		if len(tagUsersBytes) > 0 {
			if err = json.Unmarshal(tagUsersBytes, &tagUsers); err != nil {
				return nil, err
			}
		}
		for _, tagUser := range tagUsers {
			if tagUser == userID {
				return nil, nil
			}
		}
		tagUsers = append(tagUsers, userID)
		tagUsersBytes, err = json.Marshal(tagUsers)
		if err != nil {
			return nil, err
		}
		tr.Set(s.TagUsersSubspace.Sub(tag), tagUsersBytes)
		log.Println(userTags)
		log.Println(tagUsers)
		return nil, nil
	})
	return err
}

func (s *NotificationsStorageFDB) DeleteUserTag(userID, tag string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		user, err := s.GetUser(userID)
		if err != nil {
			return nil, err
		} else if user.AccountID != userID {
			return nil, errors.New(fmt.Sprintf("no such user %s", userID))
		}

		userTagsBytes, err := tr.Get(s.UserTagsSubspace.Sub(userID)).Get()
		if err != nil {
			return nil, err
		}
		userTags := []string{}
		if len(userTagsBytes) > 0 {
			if err := json.Unmarshal(userTagsBytes, &userTags); err != nil {
				return nil, err
			}
		}
		newUserTags := []string{}
		for _, userTag := range userTags {
			if userTag != tag {
				newUserTags = append(newUserTags, userTag)
			}
		}
		if len(newUserTags) == len(userTags) {
			return nil, nil
		}
		userTagsBytes, err = json.Marshal(newUserTags)
		if err != nil {
			return nil, err
		}
		tr.Set(s.UserTagsSubspace.Sub(userID), userTagsBytes)

		tagUsersBytes, err := tr.Get(s.TagUsersSubspace.Sub(tag)).Get()
		if err != nil {
			return nil, err
		}
		tagUsers := []string{}
		if len(tagUsersBytes) > 0 {
			if err = json.Unmarshal(tagUsersBytes, &tagUsers); err != nil {
				return nil, err
			}
		}
		newTagUsers := []string{}
		for _, tagUser := range tagUsers {
			if tagUser != userID {
				newTagUsers = append(newTagUsers, tagUser)
			}
		}
		tagUsersBytes, err = json.Marshal(newTagUsers)
		if err != nil {
			return nil, err
		}
		tr.Set(s.TagUsersSubspace.Sub(tag), tagUsersBytes)
		log.Println(newUserTags)
		log.Println(newTagUsers)
		return nil, nil
	})
	return err
}

func (s *NotificationsStorageFDB) GetUsersByTag(tag string) ([]*pb.User, error) {
	userIDs := []string{}
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		tagUsersBytes, err := tr.Get(s.TagUsersSubspace.Sub(tag)).Get()
		if err != nil {
			return nil, err
		}
		if len(tagUsersBytes) > 0 {
			if err = json.Unmarshal(tagUsersBytes, &userIDs); err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	users := []*pb.User{}
	for _, userID := range userIDs {
		user, err := s.GetUser(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (s *NotificationsStorageFDB) AddDevice(info *pb.DeviceInfo, userID string) (string, error) {
	device := &pb.Device{
		AccountID:  userID,
		DeviceID:   uuid.New().String(),
		DeviceInfo: info,
	}
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		deviceBytes, err := proto.Marshal(device)
		if err != nil {
			return nil, err
		}
		tr.Set(s.DevicesSubspace.Sub(device.DeviceID), deviceBytes)
		user, err := s.GetUser(userID)
		if err != nil {
			return nil, err
		}
		user.Devices = append(user.Devices, device)
		userBytes, err := proto.Marshal(user)
		if err != nil {
			return nil, err
		}
		tr.Clear(s.UsersSubspace.Sub(user.AccountID))
		tr.Set(s.UsersSubspace.Sub(user.AccountID), userBytes)
		return nil, nil
	})
	return device.DeviceID, err
}

func (s *NotificationsStorageFDB) GetDevice(deviceID string) (*pb.Device, error) {
	device := new(pb.Device)
	_, err := s.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		deviceBytes, err := tr.Get(s.DevicesSubspace.Sub(deviceID)).Get()
		if err != nil {
			return nil, err
		}
		err = proto.Unmarshal(deviceBytes, device)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return device, err
}

func (s *NotificationsStorageFDB) DeleteDevice(deviceID string) error {
	_, err := s.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		device, err := s.GetDevice(deviceID)
		if err != nil {
			return nil, err
		}
		if device == nil || device.AccountID == "" {
			return nil, nil
		}
		user, err := s.GetUser(device.AccountID)
		if err != nil {
			return nil, err
		}
		newDevices := make([]*pb.Device, 0)
		for _, d := range user.Devices {
			if d.DeviceID != deviceID {
				newDevices = append(newDevices, d)
			}
		}
		user.Devices = newDevices
		userBytes, err := proto.Marshal(user)
		if err != nil {
			return nil, err
		}
		tr.Clear(s.UsersSubspace.Sub(user.AccountID))
		tr.Set(s.UsersSubspace.Sub(user.AccountID), userBytes)
		tr.Clear(s.DevicesSubspace.Sub(deviceID))
		return nil, nil
	})
	return err
}
