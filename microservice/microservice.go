package microservice

import "log"
import pb "github.com/nostressdev/notifications/proto"

type service struct {
	*Config
}

func New(config *Config) pb.NotificationsServer {
	if config == nil {
		log.Fatalln("config must not be nil")
	}
	return &service{
		config,
	}
}
