package fetch

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	ErrNoValidRetryStrategy = errors.New("no valid retry strategy")
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

func New(options *Options) *Client {
	var fetch Client
	if options == nil {
		return setDefaultFetch()
	}
	if options.WithRetry {
		fetch.RetryStrategy = setDefaultRetryStrategy()
	}
	fetch.DefaultHeaders = options.DefaultHeaders
	fetch.Client = setDefaultClient()
	return &fetch
}

func (a *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return call(url, http.MethodGet, bytes.NewReader(nil), a.Client, headers, a.DefaultHeaders)
	}
	return callWithRetry(url, http.MethodGet, bytes.NewReader(nil), a.Client, a.RetryStrategy, headers, a.DefaultHeaders)
}

func (a *Client) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return call(url, http.MethodPost, body, a.Client, headers, a.DefaultHeaders)
	}
	return callWithRetry(url, http.MethodPost, body, a.Client, a.RetryStrategy, headers, a.DefaultHeaders)
}

func (a *Client) Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return call(url, http.MethodPut, body, a.Client, headers, a.DefaultHeaders)
	}
	return callWithRetry(url, http.MethodPut, body, a.Client, a.RetryStrategy, headers, a.DefaultHeaders)
}

func (a *Client) Delete(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return call(url, http.MethodDelete, body, a.Client, headers, a.DefaultHeaders)
	}
	return callWithRetry(url, http.MethodDelete, body, a.Client, a.RetryStrategy, headers, a.DefaultHeaders)
}

func (a *Client) Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return call(url, http.MethodPatch, body, a.Client, headers, a.DefaultHeaders)
	}
	return callWithRetry(url, http.MethodPatch, body, a.Client, a.RetryStrategy, headers, a.DefaultHeaders)
}

// call - creates a new HTTP request and returns an HTTP response
func call(url string, method string, body io.Reader, client httpClient, headers ...map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Response{}, err
	}
	allHeaders := mergeHeaders(headers...)
	for key, value := range allHeaders {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode > 399 {
		return resp, &APIError{
			StatusCode: resp.StatusCode,
			StatusText: http.StatusText(resp.StatusCode),
		}
	}
	return resp, err
}

// callWithRetry - wrap the call method with the retry strategy
func callWithRetry(url string, method string, body io.Reader, client httpClient, retryStrategy []time.Duration, headers ...map[string]string) (*http.Response, error) {
	logPrefix := "fetch: callWithRetry"
	var resp *http.Response
	var err error

	if len(retryStrategy) == 0 {
		return resp, ErrNoValidRetryStrategy
	}

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)

	go func() {
		for _, retryWait := range retryStrategy {
			resp, err = call(url, method, body, client, headers...)
			if err == nil || !isRecoverable(err) {
				break
			}

			fmt.Printf("%s: http %s request error [%s], will retry in [%s]", logPrefix, method, err, retryWait)
			time.Sleep(retryWait)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()

	return resp, err
}

// mergeHeaders - merge a slice of headers
func mergeHeaders(headersList ...map[string]string) map[string]string {
	mergedHeaders := map[string]string{}
	for _, headers := range headersList {
		for key, value := range headers {
			mergedHeaders[key] = value
		}
	}
	return mergedHeaders
}

// isRecoverable - checks if the response status code not within the 4XX range or is non standard status code
func isRecoverable(err error) bool {
	var apiError *APIError
	if !errors.As(err, &apiError) {
		return false
	}

	switch {
	case apiError.StatusCode == http.StatusNotImplemented:
		return false
	case (apiError.StatusCode >= http.StatusBadRequest && apiError.StatusCode < http.StatusInternalServerError):
		return false
	case apiError.StatusCode == 0, apiError.StatusCode > 599:
		return false
	}

	return true
}
