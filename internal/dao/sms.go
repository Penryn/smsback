package dao

import (
	"context"
	"smsback/internal/model"
)

func (d *Dao) SaveSmsUp(ctx context.Context, sms model.SmsUp) error {
	return d.orm.WithContext(ctx).Create(&sms).Error
}