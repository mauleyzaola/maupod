package helpers

import (
	"errors"
	"log"
	"time"
)

func RetryFunc(legend string, retries int, delay time.Duration, fn func(retry int) bool) (ok bool, err error) {
	if fn == nil {
		err = errors.New("[WARNING] missing parameter: fn")
		return
	}
	for i := 0; i < retries; i++ {
		retry := i + 1
		log.Printf("%s [%d/%d]\n", legend, retry, retries)
		if fn(retry) {
			return true, nil
		}
		time.Sleep(delay)
	}
	return false, nil
}
