package selecturl

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	timeURL1 := CheckSpeed(url1)
	timeURL2 := CheckSpeed(url2)
	if timeURL1 > timeURL2 {
		return url2
	}
	return url1
}
func CheckSpeed(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer2(url1, url2 string) string {
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}
func Racer3(url1, url2 string) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("time out")
	}
}
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
