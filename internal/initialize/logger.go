package initialize

import (
	"github.com/nk-hung/go-ecommerce-backend-api/global"
	"github.com/nk-hung/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
