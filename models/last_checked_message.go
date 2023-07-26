package models

import (
	"time"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type LastCheckedMessage struct {
	gorm.Model `json:"-"`
	channel    Chat       `json:"channel" gorm:"foreignKey:ChannelID;references:ChannelID`
	MessageID  int64      `json:"message_id" gorm:"type:int"`
	CreatedAt  *time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"type:datetime" json:"updated_at"`
}

func (lcm *LastCheckedMessage) GetLastMessageID(c Chat) (mid int64) {
	DB.Where("channel = ?", c.ID).First(&lcm)

	mid = lcm.MessageID
	return
}

func (lcm *LastCheckedMessage) CreateOrUpdate(chat *tele.Chat, mid int64) (updated bool, err error) {
	c := Chat{}

	_, err = c.GetOrCreate(chat)

	if err != nil {
		logrus.Errorf("last checked message database error: %+v", err)
		updated = false
		return
	}

	DB.Where("channel_id = ? AND message_id = ?", c.ID, mid).First(&lcm)

	now := time.Now()

	if lcm.ID != 0 {
		logrus.Warningf("message already existed skiped updating or creating! message_id: %d chat_username: %s chat_id: %d", mid, chat.Username, chat.ID)
		updated = false
		return
	}

	lcm.MessageID = mid

	lcm.UpdatedAt = &now

	err = DB.Save(lcm).Error
	return
}
