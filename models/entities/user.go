package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id" gorm:"column:user_id;primaryKey"`
	Username  string `json:"username" gorm:"column:username;size:255;unique"`
	Email     string `json:"email" gorm:"column:email;size:255;unique"`
	Password  uuid.UUID `json:"password" gorm:"column:password;size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}
