package notifications

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

type ConfigFDB struct {
	DB       fdb.Database
	Subspace subspace.Subspace
}

type implFDB struct {
	*ConfigFDB
	UsersSubspace    subspace.Subspace
	DevicesSubspace  subspace.Subspace
	UserTagsSubspace subspace.Subspace
	TagUsersSubspace subspace.Subspace
}

func (repo *implFDB) CreateVirtualUser(accountID string) (string, error) {
	user := &pb.User{
		AccountID: accountID,
	}
	_, err := repo.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		userBytes, err := proto.Marshal(user)
		if err != nil {
			return nil, err
		}
		tr.Set(repo.UsersSubspace.Sub(user.AccountID), userBytes)
		return nil, nil
	})
	return user.AccountID, err
}

func (repo *implFDB) GetUser(AccountID string) (*pb.User, error) {
	user := new(pb.User)
	_, err := repo.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		userBytes, err := tr.Get(repo.UsersSubspace.Sub(AccountID)).Get()
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

func (repo *implFDB) AddUserTag(userID, tag string) error {
	_, err := repo.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		user, err := repo.GetUser(userID)
		if err != nil {
			return nil, err
		} else if user.AccountID != userID {
			return nil, errors.New(fmt.Sprintf("no such user %s", userID))
		}

		userTagsBytes, err := tr.Get(repo.UserTagsSubspace.Sub(userID)).Get()
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
		tr.Set(repo.UserTagsSubspace.Sub(userID), userTagsBytes)

		tagUsersBytes, err := tr.Get(repo.TagUsersSubspace.Sub(tag)).Get()
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
		tr.Set(repo.TagUsersSubspace.Sub(tag), tagUsersBytes)
		log.Println(userTags)
		log.Println(tagUsers)
		return nil, nil
	})
	return err
}

func (repo *implFDB) DeleteUserTag(userID, tag string) error {
	_, err := repo.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		user, err := repo.GetUser(userID)
		if err != nil {
			return nil, err
		} else if user.AccountID != userID {
			return nil, errors.New(fmt.Sprintf("no such user %s", userID))
		}

		userTagsBytes, err := tr.Get(repo.UserTagsSubspace.Sub(userID)).Get()
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
		tr.Set(repo.UserTagsSubspace.Sub(userID), userTagsBytes)

		tagUsersBytes, err := tr.Get(repo.TagUsersSubspace.Sub(tag)).Get()
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
		tr.Set(repo.TagUsersSubspace.Sub(tag), tagUsersBytes)
		log.Println(newUserTags)
		log.Println(newTagUsers)
		return nil, nil
	})
	return err
}

func (repo *implFDB) GetUsersByTag(tag string) ([]*pb.User, error) {
	userIDs := []string{}
	_, err := repo.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		tagUsersBytes, err := tr.Get(repo.TagUsersSubspace.Sub(tag)).Get()
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
		user, err := repo.GetUser(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (repo *implFDB) AddDevice(info *pb.DeviceInfo, userID string) (string, error) {
	device := &pb.Device{
		AccountID:  userID,
		DeviceID:   uuid.New().String(),
		DeviceInfo: info,
	}
	_, err := repo.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		deviceBytes, err := proto.Marshal(device)
		if err != nil {
			return nil, err
		}
		tr.Set(repo.DevicesSubspace.Sub(device.DeviceID), deviceBytes)
		user, err := repo.GetUser(userID)
		if err != nil {
			return nil, err
		}
		user.Devices = append(user.Devices, device)
		userBytes, err := proto.Marshal(user)
		if err != nil {
			return nil, err
		}
		tr.Clear(repo.UsersSubspace.Sub(user.AccountID))
		tr.Set(repo.UsersSubspace.Sub(user.AccountID), userBytes)
		return nil, nil
	})
	return device.DeviceID, err
}

func (repo *implFDB) GetDevice(deviceID string) (*pb.Device, error) {
	device := new(pb.Device)
	_, err := repo.DB.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		deviceBytes, err := tr.Get(repo.DevicesSubspace.Sub(deviceID)).Get()
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

func (repo *implFDB) DeleteDevice(deviceID string) error {
	_, err := repo.DB.Transact(func(tr fdb.Transaction) (interface{}, error) {
		device, err := repo.GetDevice(deviceID)
		if err != nil {
			return nil, err
		}
		if device == nil || device.AccountID == "" {
			return nil, nil
		}
		user, err := repo.GetUser(device.AccountID)
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
		tr.Clear(repo.UsersSubspace.Sub(user.AccountID))
		tr.Set(repo.UsersSubspace.Sub(user.AccountID), userBytes)
		tr.Clear(repo.DevicesSubspace.Sub(deviceID))
		return nil, nil
	})
	return err
}
