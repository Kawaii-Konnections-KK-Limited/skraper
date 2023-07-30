package models

import (
	"fmt"
	"time"

	"github.com/gotd/td/tg"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model      `json:"-"`
	ChannelUsername *string    `gorm:"index;type:varchar(35)" json:"username"`
	ChannelID       int64      `gorm:"index" json:"channel_id"`
	CreatedAt       *time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"type:datetime" json:"updated_at"`
}

func (c *Channel) GetOrCreate(chat *tg.Channel) (created bool, err error) {
	fmt.Println(c)
	DB.Where("channel_id = ?", chat.ID).First(c)

	if c.ID != 0 {
		created = false
		return
	}

	now := time.Now()

	c.ChannelID = chat.ID
	c.ChannelUsername = &chat.Username
	c.CreatedAt = &now
	c.UpdatedAt = &now

	err = DB.Create(c).Error
	return
}
