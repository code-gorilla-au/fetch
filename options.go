package fetch

import (
	"net"
	"net/http"
	"time"
)

type Options struct {
	// Use library's default retry strategy. Default is true
	WithRetry bool
	// set headers to be send out with every request, default is none
	DefaultHeaders map[string]string
	// Provide custom retry strategy
	RetryStrategy *[]time.Duration
	// Provide a custom retry strategy
	HTTPClient *http.Client
}

type FnOpts = func(o *Options) error

// WithOpts - use functional options to provide additional config
func WithOpts(opts ...FnOpts) *Options {
	o := Options{}

	for _, fn := range opts {
		if fn == nil {
			continue
		}
		err := fn(&o)
		if err != nil {
			panic("Error setting functional option")
		}
	}

	return &o
}

// WithRetryStrategy - set custom retry strategy
func WithRetryStrategy(strategy *[]time.Duration) FnOpts {
	return func(o *Options) error {
		o.RetryStrategy = strategy
		return nil
	}
}

// WithHTTPClient - set custom http client
func WithHTTPClient(client *http.Client) FnOpts {
	return func(o *Options) error {
		o.HTTPClient = client
		return nil
	}
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

func setDefaultFetch() *Client {
	return &Client{
		RetryStrategy: setDefaultRetryStrategy(),
		Client:        setDefaultClient(),
	}
}
