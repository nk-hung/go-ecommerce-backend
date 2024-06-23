package initialize

import (
	"fmt"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
