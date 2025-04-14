package selecturl

import (
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
