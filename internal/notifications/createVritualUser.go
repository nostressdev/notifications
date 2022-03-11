package notifications

import (
	"context"
	"fmt"

	"github.com/nostressdev/nerrors"
	pb "github.com/nostressdev/notifications/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *NotificationsService) CreateVirtualUser(ctx context.Context, request *pb.CreateVirtualUserRequest) (*pb.CreateVirtualUserResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, nerrors.BadRequest.Wrap(err, "validate request")
	}
	if payload, err := s.VerifyContext(ctx, ""); err != nil {
		return nil, err
	} else {
		accountID := request.AccountID
		if request.AccountID == "" {
			rawAccountID, ok := payload["account_id"]
			if !ok {
				return nil, nerrors.BadRequest.New("wrong metadata")
			}
			accountID, ok = rawAccountID.(string)
			if !ok {
				return nil, nerrors.BadRequest.New("wrong metadata")
			}
		}
		userID, err := s.Storage.CreateVirtualUser(accountID)
		if err != nil {
			return nil, err
		}
		response := &pb.CreateVirtualUserResponse{
			AccountID: userID,
		}
		if err := response.Validate(); err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to validate response %v", err.Error()))
		}
		return response, nil
	}
}
