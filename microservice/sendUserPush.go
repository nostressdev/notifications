package microservice

import (
	"context"
	"fmt"

	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) SendUserPush(ctx context.Context, request *pb.SendUserPushRequest) (*pb.SendUserPushResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate request %v", err.Error()))
	}
	user, err := s.Repository.GetUser(request.AccountID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, device := range user.Devices {
		request := &pb.SendDevicePushRequest{
			Notification: request.Notification,
		}
		if device.DeviceInfo.DeviceType == pb.DeviceType_HUAWEI {
			_, err = s.HuaweiApp.SendMessage(request, device)
		} else if device.DeviceInfo.DeviceType == pb.DeviceType_EMAIL {
			_, err = s.EmailApp.SendMessage(request, device)
		} else {
			_, err = s.FirebaseApp.SendMessage(request, device)
		}
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to send message: %v", err.Error()))
		}
	}
	response := &pb.SendUserPushResponse{}
	if err := response.Validate(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response %v", err.Error()))
	}
	return response, nil
}
