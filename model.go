package main

import "sync"

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

// map to store the previously used Session Keys
var mKeys = make(uniqueSessionKeysType, 10)

// mutex to let multiple threads work with the mKeys map above safely
var mu = sync.Mutex{}
