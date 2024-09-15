package service

import (
	"context"
	"smsback/internal/dao"


	"gorm.io/gorm"
)

var (
	ctx = context.Background()
	d  *dao.Dao
)


func ServiceInit(db *gorm.DB){
	d=dao.New(db)
}


