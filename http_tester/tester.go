package http_tester

import (
	"loadTester/types"
	"log"
	"net/http"
	"time"
)

type EndpointInfo struct {
	Endpoint string     `json:"endpoint,omitempty"`
	Method   string     `json:"method,omitempty"`
	Headers  [][]string `json:"headers,omitempty"`
}

func (e *EndpointInfo) TestEndpoint() types.Result {
	var err error
	var resp *http.Response
	start := time.Now()
	switch e.Method {
	case "get":
		resp, err = http.Get(e.Endpoint)
	}
	if err != nil {
		log.Println(err)
	}
	finish := time.Since(start).String()
	if resp.StatusCode == http.StatusOK {
		return types.Result{
			Connectable: true,
			Duration:    finish,
		}
	}
	return types.Result{
		Connectable: false,
		Duration:    finish,
	}
}
