package controller

import (
	"MobieFund/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	contentLength := c.Request.ContentLength
	var buf = make([]byte, contentLength)
	n, _ := c.Request.Body.Read(buf)

	//surveyJson := string(buf[0:n])

	var user model.User
	json.Unmarshal(buf[0:n], &user)

	if err := user.Create(); err != nil {
		c.JSON(400, gin.H{
			"result": err,
		})

	} else {
		c.JSON(200, gin.H{
			"result": "创建成功",
		})
	}

}
