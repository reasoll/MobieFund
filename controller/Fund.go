package controller

import (
	"MobieFund/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func AddFund(c *gin.Context) {

	contentLength := c.Request.ContentLength
	var buf = make([]byte, contentLength)
	n, _ := c.Request.Body.Read(buf)

	//surveyJson := string(buf[0:n])

	var fund model.Fund
	json.Unmarshal(buf[0:n], &fund)

	fund.Id = 0
	if err := fund.Add(); err != nil {

		c.JSON(400, gin.H{
			"result": err,
		})

	} else {
		c.JSON(200, gin.H{
			"result": "创建成功",
		})
	}

}
