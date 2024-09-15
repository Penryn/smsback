package database

import (
	"smsback/internal/model"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.SmsUp{},
	)
}
