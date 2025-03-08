package main

import (
	"bytes"
	"encoding/json"
	"io"
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
	//resp, err := http.Post("http://localhost:8080/http_test", "application/json")
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("received something different than an http 200")
	}
	r := "{\"Connectable\":true}"
	if string(body) != r {
		t.Errorf("Unexpected body response got %q wanted %v", string(body), r)
	}

}
