package microservice

import (
	"context"
	"fmt"

	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DeleteUserTag(ctx context.Context, request *pb.DeleteUserTagRequest) (*pb.DeleteUserTagResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate request %v", err.Error()))
	}
	err := s.Repository.DeleteUserTag(request.AccountID, request.Tag)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.DeleteUserTagResponse{}
	if err := response.Validate(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response %v", err.Error()))
	}
	return response, nil
}
