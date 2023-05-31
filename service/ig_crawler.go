package service

import (
	"net/http"
	"trahsu-tools/models"

	"github.com/gin-gonic/gin"
)

func Crawler(c *gin.Context) {

	userId, valid := c.GetQuery("userId")
	if !valid {
		c.JSON(http.StatusOK, gin.H{
			"userId": "Not Found",
		})
	}

	imgSrcList := models.GetPublicPostHandler(c.Writer, c.Request)

	c.JSON(http.StatusOK, gin.H{
		"userId":    userId,
		"filesName": imgSrcList,
	})

}
