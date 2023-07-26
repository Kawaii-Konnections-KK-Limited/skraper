package models

import (
	"time"

	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model      `json:"-"`
	ChannelUsername *string    `gorm:"index;type:varchar(35)" json:"username"`
	ChannelID       int64      `gorm:"index" json:"user_id"`
	CreatedAt       *time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"type:datetime" json:"updated_at"`
}

func (c *Chat) GetOrCreate(chat *tele.Chat) (created bool, err error) {
	DB.Where("channel_id = ?", c.ChannelID).First(&c)

	if c.ID != 0 {
		created = false
		return
	}

	now := time.Now()

	c.ChannelID = chat.ID
	c.ChannelUsername = &chat.Username
	c.CreatedAt = &now
	c.UpdatedAt = &now

	return
}
