package migrations

import (
	"log/slog"

	"github.com/enesakbal/go-chat-server/models/entities"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Chat{},
		&entities.Message{},
	)

	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	slog.Info("Database Migrated")


}

func RunMigrationDrop(db *gorm.DB) {
	err := db.Migrator().DropTable(&entities.User{},
		&entities.Message{},
		&entities.Chat{},
	)

	if err != nil {
		slog.Error("Migration drop failed: " + err.Error())
		panic("Migration drop failed: " + err.Error())
	}

	err = db.AutoMigrate(&entities.User{}, &entities.Message{}, &entities.Chat{})
	if err != nil {
		slog.Error("AutoMigrate failed after drop: " + err.Error())
		panic("AutoMigrate failed after drop: " + err.Error())	
	}

	slog.Info("Database Fresh")
}
