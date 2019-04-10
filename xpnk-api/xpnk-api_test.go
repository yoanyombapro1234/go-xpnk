package main

import (
		"testing"
		"net/http"
  		//"net/http/httptest"
 )

func TestV2Ping(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v2/ping", nil)
	if err != nil {
		t.Fatal(err)
	} else if req.Code != 200 {
		t.Errorf("error %v", req.Code)
	}
}