package database

import (
	enviroment "github.com/enesakbal/go-chat-server/bootstrap/enviroment"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseInit(env *enviroment.Enviroment) *gorm.DB {
	host := env.DBHost
	port := env.DBPort
	username := env.DBUser
	password := env.DBPass
	database := env.DBName

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}
