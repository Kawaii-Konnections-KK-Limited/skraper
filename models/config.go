package models

import (
	"time"

	"github.com/gotd/td/tg"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model `json:"-"`
	ChannelID  int64      `gorm:"index" json:"user_id"`
	Link       string     `gorm:"type:string;unique" json:"link"`
	CreatedAt  *time.Time `json:"last_activity" gorm:"type:datetime"`
}

func (l *Link) Create(chat *tg.Channel, link string) (created bool, err error) {
	c := Channel{}

	_, err = c.GetOrCreate(chat)

	if err != nil {
		created = false
		logrus.Errorf("error caused while creating the chat (%s): %s", chat.Title, err)
		return
	}

	DB.Where("link = ?", link).First(l)

	if l.ID != 0 {
		created = false
		logrus.Infof("record already existed [channel_username:%s link:%s]", *c.ChannelUsername, link)
		return
	}

	now := time.Now()

	l.ChannelID = int64(c.ID)
	l.CreatedAt = &now
	l.Link = link

	err = DB.Create(l).Error
	if err == nil {
		logrus.Infof("created Link row with values [channel_username:%s link:%s]", *c.ChannelUsername, link)
	}

	return
}
