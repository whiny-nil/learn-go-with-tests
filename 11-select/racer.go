package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureResponseTime(a)
	durationB := measureResponseTime(b)

	if durationA <= durationB {
		return a
	} else {
		return b
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}
