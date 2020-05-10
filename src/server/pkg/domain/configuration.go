package domain

import (
	"time"
)

type Configuration struct {
	Stores  []FileStore
	PgConn  string
	Retries int
	Delay   time.Duration
	Port    string
}

func (c *Configuration) Validate() error {
	for _, v := range c.Stores {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
