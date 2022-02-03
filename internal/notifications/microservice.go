package notifications

import (
	"log"

	"github.com/nostressdev/notifications/internal/storage"
	pb "github.com/nostressdev/notifications/proto"
	"github.com/nostressdev/notifications/utils"
	"github.com/nostressdev/signer"
)

type Config struct {
	*signer.Signer
	Storage     storage.NotificationsStorage
	FirebaseApp *utils.FirebaseApp
	HuaweiApp   *utils.HuaweiApp
	EmailApp    *utils.EmailApp
}

type NotificationsService struct {
	*Config
}

func New(config *Config) pb.NotificationsServer {
	if config == nil {
		log.Fatalln("config must not be nil")
	}
	return &NotificationsService{
		config,
	}
}
