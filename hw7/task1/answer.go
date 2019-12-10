package main

import "net/http"

type Answer struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestUri string      `json:"request_uri"`
	Header     http.Header `json:"headers"`
}
