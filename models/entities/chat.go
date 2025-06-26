package entities

import "time"

type Chat struct {
	ID        string    `json:"id" gorm:"column:chat_id;primaryKey"`
	OwnerID   string    `json:"owner_id" gorm:"column:owner_id"`
	Title     string    `json:"title" gorm:"column:title;size:255"`
	Messages  []Message `json:"messages" gorm:"foreignKey:ChatID;"`
	IsPublic  bool      `json:"is_public" gorm:"column:is_public;default:false"`
	IsDeleted bool      `json:"is_deleted" gorm:"column:is_deleted;default:false"`
	IsPinned  bool      `json:"is_pinned" gorm:"column:is_pinned;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}