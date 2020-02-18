package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InputPayload represents the expected object posted to the /isgood path
type InputPayload struct {
	CheckType       string `form:"checkType" json:"checkType" binding:"required"`
	ActivityType    string `form:"activityType" json:"activityType" binding:"required"`
	CheckSessionKey string `form:"checkSessionKey" json:"checkSessionKey" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/isgood", func(c *gin.Context) {
		var json InputPayload

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, json)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
