package util

import (
	"github.com/gin-gonic/gin"
)

func RespOK(data any) gin.H {
	return gin.H{
		"success": true,
		"msg":     "",
		"data":    data,
	}
}

func RespMsg(msg string) gin.H {
	return gin.H{
		"success": true,
		"msg":     msg,
		"data":    nil,
	}
}

func RespFail(msg string) gin.H {
	return gin.H{
		"success": false,
		"msg":     msg,
		"data":    nil,
	}
}
