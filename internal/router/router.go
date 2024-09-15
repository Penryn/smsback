package router

import (
	"smsback/internal/handler/sms"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/sms/up", sms.HandleSmsUp)

	}
}
