package helpers

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

func RetryFunc(legend string, fn func(retry int) bool) (ok bool, err error) {
	maxRetries := viper.GetInt("max-retries")
	if maxRetries == 0 {
		err = errors.New("cannot resolve variable: max-retries")
		return
	}
	delay := viper.GetDuration("DELAY")
	if delay == 0 {
		err = errors.New("cannot resolve variable: DELAY")
		return
	}
	if fn == nil {
		err = errors.New("[WARNING] missing parameter: fn")
		return
	}
	for i := 0; i < maxRetries; i++ {
		retry := i + 1
		log.Printf("%s [%d/%d]\n", legend, retry, maxRetries)
		if fn(retry) {
			return true, nil
		}
		time.Sleep(delay)
	}
	return false, nil
}
