package models

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model `json:"-"`
	ChannelID  int64      `gorm:"index" json:"user_id"`
	CreatedAt  *time.Time `json:"last_activity" gorm:"type:datetime"`
}
