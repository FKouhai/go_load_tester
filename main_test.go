package main

import (
	"bytes"
	"encoding/json"
	"loadTester/http_tester"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	ep := http_tester.EndpointInfo{
		Endpoint: "http://localhost:8080/",
		Method:   "get",
	}
	jd, err := json.Marshal(ep)
	if err != nil {
		t.Fail()
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/http_test", bytes.NewBuffer(jd))
	if err != nil {
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fail()
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("wanted %q got %v", http.StatusOK, resp.StatusCode)
	}

}
