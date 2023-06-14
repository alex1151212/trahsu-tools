package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PunchIn(c *gin.Context) {

	userId, valid := c.GetQuery("userId")
	if !valid {
		c.JSON(http.StatusOK, gin.H{
			"userId": "Not Found",
		})
	}
	//TODO 將打卡時間存入資料庫
	//TODO 判斷今天是否打卡過，若有則忽略

	c.JSON(http.StatusOK, gin.H{
		"userId":      userId,
		"punchInTime": "123",
		"offWorkTime": "123",
	})

}
