package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMalformedJSONValidation(t *testing.T) {
	// Arrange
	// Prepare a corrupted JSON as input
	var body = `{ "checkType": , }`
	// Create a new server, set up a handler and prepare a request to be sent
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/isgood", IsGoodHandler)
	req, _ := http.NewRequest("POST", "/isgood", strings.NewReader(body))

	// Act
	// Send the request to the server
	r.ServeHTTP(w, req)

	// Assert
	// Verify the received http status code
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}

	// Convert the JSON response to a map
	var response map[string]string
	_ = json.Unmarshal([]byte(w.Body.String()), &response)

	// Take the response value and check it
	msg, exists := response["error"]
	if !exists || !strings.HasPrefix(msg, "invalid character") {
		t.Fail()
	}
}

type TestCase struct {
	FieldToCheck string
	Body         string
	IsError      bool
}

func TestCheckFieldsValidation(t *testing.T) {
	// Arrange
	// Prepare test cases with an incorrect JSON value for one of the 'enum' fields
	cases := []TestCase{
		TestCase{
			FieldToCheck: "CheckType",
			Body: `{
				"checkType": "UNKNOWN",
				"activityType": "SIGNUP",
				"checkSessionKey": "1",
				"activityData": [{
					"kvpKey": "key1",
					"kvpValue": "value1",
					"kvpType": "general.string"
				}]
			}`,
			IsError: true,
		},
		TestCase{
			FieldToCheck: "ActivityType",
			Body: `{
				"checkType": "BIOMETRIC",
				"activityType": "UNKNOWN",
				"checkSessionKey": "2",
				"activityData": [{
					"kvpKey": "key1",
					"kvpValue": "value1",
					"kvpType": "general.string"
				}]
			}`,
			IsError: true,
		},
		TestCase{
			FieldToCheck: "KvpType",
			Body: `{
				"checkType": "COMBO",
				"activityType": "PAYMENT",
				"checkSessionKey": "3",
				"activityData": [{
					"kvpKey": "key1",
					"kvpValue": "value1",
					"kvpType": "UNKNOWN"
				}]
			}`,
			IsError: true,
		},
	}

	for _, item := range cases {

		// Create a new server, set up a handler and prepare a request to be sent
		w := httptest.NewRecorder()
		r := gin.Default()

		r.POST("/isgood", IsGoodHandler)
		req, _ := http.NewRequest("POST", "/isgood", strings.NewReader(item.Body))

		// Act
		// Send the request to the server
		r.ServeHTTP(w, req)

		// Assert
		// Verify the received http status code
		if w.Code != http.StatusBadRequest {
			t.Fail()
		}

		// Convert the JSON response to a map
		var response map[string]string
		_ = json.Unmarshal([]byte(w.Body.String()), &response)

		// Take the response value and check it
		msg, exists := response["error"]
		switch item.FieldToCheck {
		case "CheckType":
			if !exists || !strings.HasPrefix(msg, "Key: 'InputPayload.CheckType' Error:Field validation for 'CheckType' failed on the 'oneof' tag") {
				t.Fail()
			}
		case "ActivityType":
			if !exists || !strings.HasPrefix(msg, "Key: 'InputPayload.ActivityType' Error:Field validation for 'ActivityType' failed on the 'oneof' tag") {
				t.Fail()
			}
		case "KvpType":
			if !exists || !strings.HasPrefix(msg, "Key: 'InputPayload.ActivityData[0].KvpType' Error:Field validation for 'KvpType' failed on the 'oneof' tag") {
				t.Fail()
			}
		}
	}
}

func TestSessionKeyValidation(t *testing.T) {
	// Arrange
	// Prepare two requests containing the same Session Key
	var body1 = `{
		"checkType": "DEVICE",
		"activityType": "SIGNUP",
		"checkSessionKey": "1234",
		"activityData": [{
			"kvpKey": "key1",
			"kvpValue": "value1",
			"kvpType": "general.string"
		}]
	}`
	var body2 = `{
		"checkType": "DEVICE",
		"activityType": "SIGNUP",
		"checkSessionKey": "1234",
		"activityData": [{
			"kvpKey": "key1",
			"kvpValue": "value1",
			"kvpType": "general.string"
		}]
	}`
	// Create a new server, set up a handler and prepare a request to be sent
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/isgood", IsGoodHandler)
	req1, _ := http.NewRequest("POST", "/isgood", strings.NewReader(body1))
	req2, _ := http.NewRequest("POST", "/isgood", strings.NewReader(body2))

	// Act
	// Send the request to the server
	r.ServeHTTP(w1, req1)
	r.ServeHTTP(w2, req2)

	// Assert
	// Verify that the received http status code is 200 for the first request
	if w1.Code != http.StatusOK {
		t.Fail()
	}

	// Verify the received http status code is 400 for the second request
	if w2.Code != http.StatusBadRequest {
		t.Fail()
	}

	// Convert the JSON response to a map
	var response map[string]string
	_ = json.Unmarshal([]byte(w2.Body.String()), &response)

	// Take the response value and check it
	msg, exists := response["error"]
	if !exists || !strings.HasPrefix(msg, "Session Key has already been used before") {
		t.Fail()
	}
}

func TestPathNotFound(t *testing.T) {
	// Arrange
	// Create a new server and send a request to a non-existent /ping path
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/isgood", IsGoodHandler)
	req, _ := http.NewRequest("POST", "/ping", nil)

	// Act
	// Send the request to the server
	r.ServeHTTP(w, req)

	// Assert
	// Verify the received http status code
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}
