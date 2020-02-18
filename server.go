package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InputPayload represents the expected object posted to the /isgood path
type InputPayload struct {
	CheckType       string         `json:"checkType" binding:"required,oneof=DEVICE BIOMETRIC COMBO"`
	ActivityType    string         `json:"activityType" binding:"required,oneof=SIGNUP LOGIN PAYMENT CONFIRMATION"`
	CheckSessionKey string         `json:"checkSessionKey" binding:"required"`
	ActivityData    []ActivityData `json:"activityData" binding:"required,dive"`
}

// ActivityData represents the internal object for the InputPayload type above
type ActivityData struct {
	KvpKey   string `json:"kvpKey" binding:"required"`
	KvpValue string `json:"kvpValue" binding:"required"`
	KvpType  string `json:"kvpType" binding:"required,oneof='general.string' 'general.integer' 'general.float' 'general.bool'"`
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
