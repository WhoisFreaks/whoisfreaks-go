package whoisfreaks

import (
	"math"
	"net/http"
	"time"
)

// RetryTransport wraps an http.RoundTripper and retries transient failures
// with exponential backoff. Attach it to the generated client's http.Client.
type RetryTransport struct {
	Base       http.RoundTripper
	MaxRetries int
	BaseDelay  time.Duration
}

var retryable = map[int]bool{429: true, 500: true, 502: true, 503: true, 504: true}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	max := t.MaxRetries
	if max == 0 {
		max = 3
	}
	delay := t.BaseDelay
	if delay == 0 {
		delay = 500 * time.Millisecond
	}
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= max; attempt++ {
		resp, err = base.RoundTrip(req)
		if err == nil && !retryable[resp.StatusCode] {
			return resp, nil
		}
		if attempt < max {
			time.Sleep(time.Duration(float64(delay) * math.Pow(2, float64(attempt))))
		}
	}
	return resp, err
}
