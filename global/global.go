package global

import (
	"github.com/nk-hung/go-ecommerce-backend-api/pkg/logger"
	"github.com/nk-hung/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
