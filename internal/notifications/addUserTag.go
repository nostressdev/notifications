package notifications

import (
	"context"
	"fmt"

	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *NotificationsService) AddUserTag(ctx context.Context, request *pb.AddUserTagRequest) (*pb.AddUserTagResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate request %v", err.Error()))
	}
	err := s.Storage.AddUserTag(request.AccountID, request.Tag)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.AddUserTagResponse{}
	if err := response.Validate(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response %v", err.Error()))
	}
	return response, nil
}
