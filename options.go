package fetch

import (
	"net"
	"net/http"
	"time"
)

type AxiosOptions struct {
	WithRetry      bool
	DefaultHeaders map[string]string
}

// setDefaultRetryStrategy - sets the retry attempts
func setDefaultRetryStrategy() []time.Duration {
	return []time.Duration{
		1 * time.Second,
		3 * time.Second,
		5 * time.Second,
		10 * time.Second,
	}
}

// setDefaultClient - returns the default http client
func setDefaultClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 15 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Second * 15,
	}
}

func setDefaultAxios() *Axios {
	return &Axios{
		retryStrategy: setDefaultRetryStrategy(),
		client:        setDefaultClient(),
	}
}
