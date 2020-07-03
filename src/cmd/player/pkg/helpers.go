package pkg

import "time"

func sleep(secs int) {
	time.Sleep(time.Second * time.Duration(secs))
}
