package utils

import (
	"context"
	"fmt"

	"github.com/nostressdev/signer"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func GetServerInterceptor(logger *zap.Logger, signer *signer.Signer) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		msg := req.(proto.Message)
		logger.Debug(fmt.Sprintf("Request: %s %s", info.FullMethod, protojson.Format(msg)))
		//methodPath := info.FullMethod
		//authPrefix := append(make([]interface{}, 0), interface{}("/syntoks.Auth"))
		//if !belongsToPrefixes(methodPath, authPrefix) {
		//	whitelist := incomingRequest.Claims[PropertyWhitelist].([]interface{})
		//	blacklist := incomingRequest.Claims[PropertyBlacklist].([]interface{})
		//	inWhitelist := belongsToPrefixes(methodPath, whitelist)
		//	inBlacklist := belongsToPrefixes(methodPath, blacklist)
		//	if !inWhitelist || inBlacklist {
		//		errStr := "failed to validate incoming request: permission denied"
		//		logger.Warn(errStr)
		//		return nil, status.Error(codes.PermissionDenied, errStr)
		//	}
		//}
		resp, err := handler(ctx, req)
		if err == nil && resp != nil {
			msg = resp.(proto.Message)
			logger.Debug(fmt.Sprintf("Response: %s %s", info.FullMethod, protojson.Format(msg)))
		} else if errStatus := status.Convert(err); errStatus.Code() == codes.PermissionDenied {
			logger.Warn(errStatus.Message())
		} else if errStatus.Code() == codes.InvalidArgument {
			logger.Warn(errStatus.Message())
		} else if errStatus.Code() == codes.NotFound {
			logger.Warn(errStatus.Message())
		} else {
			logger.Error(errStatus.Message())
		}
		return resp, err
	})
}
