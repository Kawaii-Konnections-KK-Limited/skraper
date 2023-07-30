package models

import (
	"time"

	"github.com/gotd/td/tg"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LastCheckedMessage struct {
	gorm.Model `json:"-"`
	Channel    Channel    `json:"channel" gorm:"foreignKey:ChannelID;references:ChannelID"`
	ChannelID  int64      `gorm:"type:int" json:"channel_id"`
	MessageID  int64      `json:"message_id" gorm:"type:int"`
	CreatedAt  *time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"type:datetime" json:"updated_at"`
}

func (lcm *LastCheckedMessage) GetLastMessageID(c Channel) (mid int64) {
	DB.Where("channel = ?", c.ID).First(&lcm)

	mid = lcm.MessageID
	return
}

func (lcm *LastCheckedMessage) CreateOrUpdate(chat *tg.Channel, mid int64) (updated bool, err error) {
	c := Channel{}

	_, err = c.GetOrCreate(chat)

	if err != nil {
		logrus.Errorf("last checked message database error: %+v", err)
		updated = false
		return
	}

	DB.Where("channel_id = ? AND message_id = ?", c.ID, mid).First(&lcm)

	now := time.Now()

	lcm.MessageID = mid
	lcm.UpdatedAt = &now

	if lcm.ID != 0 {
		logrus.Warningf("message already existed skiped updating or creating! message_id: %d chat_username: %s chat_id: %d", mid, chat.Username, chat.ID)
		updated = true
		err = DB.Save(lcm).Error
	} else {
		lcm.ChannelID = int64(c.ID)
		err = DB.Create(lcm).Error
	}

	return
}
