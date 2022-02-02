package microservice

import (
	"github.com/nostressdev/notifications/internal"
	"github.com/nostressdev/notifications/utils"
	"github.com/nostressdev/signer"
)

type Config struct {
	*signer.Signer
	Repository  notifications.Service
	FirebaseApp *utils.FirebaseApp
	HuaweiApp   *utils.HuaweiApp
	EmailApp    *utils.EmailApp
}
