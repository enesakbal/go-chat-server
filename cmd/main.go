package main

import (
	"github.com/enesakbal/go-chat-server/bootstrap/database"
	"github.com/enesakbal/go-chat-server/bootstrap/enviroment"
	"github.com/enesakbal/go-chat-server/bootstrap/migrations"
)

func main() {
	env := enviroment.EnviromentInit()
	db := database.DatabaseInit(env)


	migrations.RunMigration(db)

}