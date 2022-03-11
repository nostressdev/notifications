package notifications

import (
	"context"

	"github.com/nostressdev/notifications/proto"
)

func (s *NotificationsService) Ping(_ context.Context, _ *proto.PingRequest) (*proto.PingResponse, error) {
	response := &proto.PingResponse{}
	return response, nil
}
