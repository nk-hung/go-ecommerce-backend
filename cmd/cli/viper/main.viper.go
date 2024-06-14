package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Security struct {
		Jwt struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local") // ten file
	viper.SetConfigType("yaml")

	// read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w \n", err))
	}

	fmt.Println("Server Port:::", viper.GetInt("server.port"))
	fmt.Println("JWT KEy:::", viper.GetString("security.jwt.key"))

	// configurate structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}

	fmt.Println("Config Port::", config.Security.Jwt.Key)

	for _, db := range config.Databases {
		fmt.Print(111)
		fmt.Printf("database User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
