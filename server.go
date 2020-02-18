package main

import (
	"net/http"
	"sync"

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

type uniqueSessionKeysType map[string]bool

func main() {
	r := gin.Default()

	// map to store the previously used Session Keys
	mKeys := make(uniqueSessionKeysType, 10)
	// mutex to let multiple threads work with the mKeys map above safely
	mu := &sync.Mutex{}

	r.POST("/isgood", func(c *gin.Context) {
		var json InputPayload

		// parse JSON in automated mode according to the definition above
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check that the Session Key has not been already used
		if _, keyExist := mKeys[json.CheckSessionKey]; keyExist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Session Key has already been used before"})
			return
		}

		// lock the mutex to let only one thread to go in
		mu.Lock()
		// add a new Session Key to the map of previously used keys
		mKeys[json.CheckSessionKey] = true
		// unlock the mutex
		mu.Unlock()

		// respond with valid JSON
		c.JSON(http.StatusOK, json)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
