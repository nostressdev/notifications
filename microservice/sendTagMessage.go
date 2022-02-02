package microservice

import (
	"context"
	"fmt"

	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) SendTagPush(ctx context.Context, request *pb.SendTagPushRequest) (*pb.SendTagPushResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate request %v", err.Error()))
	}
	users, err := s.Repository.GetUsersByTag(request.Tag)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, user := range users {
		_, err := s.SendUserPush(ctx, &pb.SendUserPushRequest{
			AccountID:    user.AccountID,
			Notification: request.Notification,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to send push %v", err.Error()))
		}
	}
	response := &pb.SendTagPushResponse{}
	if err := response.Validate(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response %v", err.Error()))
	}
	return response, nil
}
