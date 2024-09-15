package service

import "smsback/internal/model"

func SaveSmsUp(sms model.SmsUp) error {
	return d.SaveSmsUp(ctx,sms)
}