package models

import (
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

const (
	CoinRate = 1000000000
)

func InitDB() (err error) {
	log.Info("Initializing database...")
	gCnf := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	uri := "db.sqlite3"
	dbType := "sqlite"
	if dbType == "mysql" {
		DB, err = gorm.Open(mysql.Open(uri), gCnf)
	} else if dbType == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(uri), gCnf)
	} else {
		log.Panic("Unknown database type: ", dbType)
	}

	if err != nil {
		log.Panic("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(
		&Channel{},
		&Link{},
		&LastCheckedMessage{},
	)
	log.Info("Database initialized")
	return
}
