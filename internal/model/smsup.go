package model

import "time"

type SmsUp struct {
	ID          uint      `gorm:"primaryKey"`
	SignName    string    `json:"sign_name"`     // 签名
	PhoneNumber string    `json:"phone_number" binding:"required"`  // 手机号
	Content     string    `json:"content" binding:"required"`       // 短信内容
	SendTime    string    `json:"send_time" binding:"required"`     // 发送时间
	DestCode    string    `json:"dest_code" binding:"required"`     // 扩展码
	SequenceID  int    `json:"sequence_id" binding:"required"`   // 序列号
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // 创建时间
}
