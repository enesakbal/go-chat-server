package entities

import "time"

type Message struct {
	ID        string `json:"id" gorm:"column:message_id;primaryKey"`
	ChatID    string `json:"chat_id" gorm:"column:chat_id;size:255"`
	OwnerID   string `json:"owner_id" gorm:"column:owner_id;size:255"`
	Content   string `json:"content" gorm:"column:content;size:255"`
	SentTime  time.Time `json:"sent_time" gorm:"column:sent_time;autoCreateTime"`
}