package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	dealtime := time.Now().Add(timeout)
	//一分钟重试
	for tries := 0; time.Now().Before(dealtime); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying...", err)
		time.Sleep(time.Second << tries)
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
