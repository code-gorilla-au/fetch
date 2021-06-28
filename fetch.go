package fetch

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
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
	retryStrategy  []time.Duration
	client         httpClient
	defaultHeaders map[string]string
}

var _ client = (*Client)(nil)

func New(options *Options) *Client {
	var fetch Client
	if options == nil {
		return setDefaultFetch()
	}
	if options.WithRetry {
		fetch.retryStrategy = setDefaultRetryStrategy()
	}
	fetch.defaultHeaders = options.DefaultHeaders
	fetch.client = setDefaultClient()
	return &fetch
}

func (a *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	if a.retryStrategy == nil {
		return call(url, http.MethodGet, bytes.NewReader(nil), a.client, headers)
	}
	return callWithRetry(url, http.MethodGet, bytes.NewReader(nil), a.client, a.retryStrategy, headers)
}

func (a *Client) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.retryStrategy == nil {
		return call(url, http.MethodPost, body, a.client, headers)
	}
	return callWithRetry(url, http.MethodPost, body, a.client, a.retryStrategy, headers)
}

func (a *Client) Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.retryStrategy == nil {
		return call(url, http.MethodPut, body, a.client, headers)
	}
	return callWithRetry(url, http.MethodPut, body, a.client, a.retryStrategy, headers)
}

func (a *Client) Delete(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.retryStrategy == nil {
		return call(url, http.MethodDelete, body, a.client, headers)
	}
	return callWithRetry(url, http.MethodDelete, body, a.client, a.retryStrategy, headers)
}

func (a *Client) Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.retryStrategy == nil {
		return call(url, http.MethodPatch, body, a.client, headers)
	}
	return callWithRetry(url, http.MethodPatch, body, a.client, a.retryStrategy, headers)
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

	for _, retryWait := range retryStrategy {
		resp, err = call(url, http.MethodPost, body, client, headers...)
		if err == nil {
			if resp.StatusCode == http.StatusTooManyRequests {
				return resp, errors.New(http.StatusText(resp.StatusCode))
			}
			return resp, nil
		}
		fmt.Printf("%s: http %s request error [%s], will retry in [%s]", logPrefix, method, err, retryWait)
		time.Sleep(retryWait)
	}

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
