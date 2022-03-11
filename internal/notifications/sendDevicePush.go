package notifications

import (
	"context"
	"fmt"

	"github.com/nostressdev/nerrors"
	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *NotificationsService) SendDevicePush(ctx context.Context, request *pb.SendDevicePushRequest) (*pb.SendDevicePushResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate request %v", err.Error()))
	}
	if payload, err := s.VerifyContext(ctx, ""); err != nil {
		return nil, err
	} else {
		rawAccountID, ok := payload["account_id"]
		if !ok {
			return nil, nerrors.BadRequest.New("wrong metadata")
		}
		accountID, ok := rawAccountID.(string)
		if !ok {
			return nil, nerrors.BadRequest.New("wrong metadata")
		}
		device, err := s.Storage.GetDevice(accountID, request.DeviceID)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if device.DeviceInfo.DeviceType == pb.DeviceType_HUAWEI {
			_, err = s.HuaweiApp.SendMessage(request, device)
		} else if device.DeviceInfo.DeviceType == pb.DeviceType_EMAIL {
			_, err = s.EmailApp.SendMessage(request, device)
		} else {
			_, err = s.FirebaseApp.SendMessage(request, device)
		}
		if err.Error() == "Requested entity was not found." {
			err := s.Storage.DeleteDevice(accountID, device.DeviceID)
			if err != nil {
				return &pb.SendDevicePushResponse{}, err
			}
		} else if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to send message: %v", err.Error()))
		}

		response := &pb.SendDevicePushResponse{}
		if err := response.Validate(); err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response: %v", err.Error()))
		}
		return response, nil
	}
}
