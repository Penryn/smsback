package middleware

import (
	"github.com/gin-gonic/gin"

)



// HandleNotFound
//
//	404处理
func HandleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"code": 404,
		"msg":  "404, page not found!",
	})
}
