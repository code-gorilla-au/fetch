package fetch

import (
	"io"
	"net/http"
	"time"
)

// client - basic http client with retry
type client interface {
	Get(url string, headers map[string]string) (resp *http.Response, err error)
	Post(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
	Put(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
	Patch(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
	Delete(url string, body io.Reader, headers map[string]string) (resp *http.Response, err error)
}

// httpClient - client interface
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}

type Client struct {
	// Retry backoff strategy.
	// Default is 1s,3s,5s,10s
	RetryStrategy []time.Duration
	// HTTP client
	Client httpClient
	// Headers to be added to each request
	DefaultHeaders map[string]string
}

var _ client = (*Client)(nil)
