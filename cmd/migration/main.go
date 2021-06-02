package main

import (
	"fmt"

	"github.com/chanhlab/go-service-template/config"
	"github.com/chanhlab/go-utils/database/mysql"
	"github.com/chanhlab/go-utils/logger"
)

// main ...
func main() {
	fmt.Printf("Migrate \n")
	config.NewConfig()
	config := config.AppConfig
	fmt.Printf("\ndebug %s \n", config.MySQL.Host)
	logger.Init(config.Logger.LogLevel, config.Logger.LogTimeFormat)
	db := mysql.GetConnection(config.MySQL.Host, config.MySQL.Port, config.MySQL.DBName, config.MySQL.Username, config.MySQL.Password, config.MySQL.MaxIDLEConnection, config.MySQL.MaxOpenConnection)

	if db == nil {
		panic("DB can not be nil")
	}

	// Create User table
	// err := db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	logger.Log.Sugar().Error(err)
	// }
}
