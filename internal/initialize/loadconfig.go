package initialize

import (
	"fmt"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local") // ten file
	viper.SetConfigType("yaml")

	// read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w \n", err))
	}

	fmt.Println("JWT KEy:::", viper.GetString("security.jwt.key"))
	fmt.Println("Server Port:::", viper.GetInt("server.port"))

	// configurate structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
}
