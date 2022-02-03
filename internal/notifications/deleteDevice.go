package notifications

import (
	"context"
	"fmt"

	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *NotificationsService) DeleteDevice(ctx context.Context, request *pb.DeleteDeviceRequest) (*pb.DeleteDeviceResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate request %v", err.Error()))
	}
	err := s.Storage.DeleteDevice(request.DeviceID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.DeleteDeviceResponse{}
	if err := response.Validate(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response %v", err.Error()))
	}
	return response, nil
}
