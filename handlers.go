package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsGoodHandler serves the /isgood path on the server
func IsGoodHandler(c *gin.Context) {
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
}
