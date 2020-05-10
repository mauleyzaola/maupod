package api

import "time"

type FileStore struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
}

type Configuration struct {
	Stores  []FileStore
	PgConn  string
	Retries int
	Delay   time.Duration
	Port    string
}
