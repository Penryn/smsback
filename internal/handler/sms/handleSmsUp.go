package sms

import (
	"bytes"
	"fmt"
	"io"
	"smsback/internal/model"
	"smsback/internal/pkg/utils"
	"smsback/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleSmsUp(c *gin.Context) {
	var smsUps []model.SmsUp
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 打印请求体
	fmt.Println(string(bodyBytes))

	// 恢复 Request.Body，以便之后还能继续读取
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// 解析请求数据
	if err := c.ShouldBindJSON(&smsUps); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 保存数据
	for _, smsUp := range smsUps {
		// 传递的时间为空时，设置默认值
		if smsUp.SendTime == "" {
			smsUp.SendTime = time.Now().Format("2006-01-02 15:04:05")
		}
		err := service.SaveSmsUp(smsUp)
		if err != nil {
			utils.JsonErrorResponse(c, 200500, "服务错误")
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
}
