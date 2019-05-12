package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"xpnk-api/users"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestUserByTwitter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/users/twitter/131547767", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	if assert.Equal(t, 200, w.Code){
		body := w.Body.String()
		fmt.Printf("\nBody.String: %+v\n", body)
		
		var payload users.XPNKUser
		err := json.Unmarshal(body, &payload)
		if err != nil {
			fmt.Printf("\nJSON: %+v\n", json.User_ID)
		}
		
		//assert.Equal(t, "pong", w.Body.String())
	}
	
}