package initialize

import (
	"fmt"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log OK!!", zap.String("ok", "success"))
	InitMySql()
	InitRedis()
	// InitKafka()

	r := InitRouter()

	r.Run(":8002")
}
