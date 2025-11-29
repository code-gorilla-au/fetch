package fetch

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func New(options *Options) *Client {

	if options == nil {
		return setDefaultFetch()
	}

	// defaults
	var fetch Client
	fetch.DefaultHeaders = options.DefaultHeaders
	fetch.Client = setDefaultClient()
	if options.WithRetry {
		fetch.RetryStrategy = setDefaultRetryStrategy()
	}

	// overrides
	if options.HTTPClient != nil {
		fetch.Client = options.HTTPClient
	}

	if options.RetryStrategy != nil {
		fetch.RetryStrategy = *options.RetryStrategy
	}

	return &fetch
}

func (a *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	return a.do(url, http.MethodGet, nil, headers)
}

func (a *Client) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(url, http.MethodPost, body, headers)
}

func (a *Client) Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(url, http.MethodPut, body, headers)
}

func (a *Client) Delete(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(url, http.MethodDelete, body, headers)
}

func (a *Client) Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(url, http.MethodPatch, body, headers)
}

// do - make http call with the provided configuration
func (a *Client) do(url string, method string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return a.call(url, method, body, headers, a.DefaultHeaders)
	}
	return a.callWithRetry(url, method, body, a.RetryStrategy, headers, a.DefaultHeaders)
}

// callWithRetry - wrap the call method with the retry strategy
func (a *Client) callWithRetry(url string, method string, body io.Reader, retryStrategy []time.Duration, headers ...map[string]string) (*http.Response, error) {
	logPrefix := "fetch: callWithRetry"
	var resp *http.Response
	var err error

	if len(retryStrategy) == 0 {
		return resp, ErrNoValidRetryStrategy
	}

	for _, retryWait := range retryStrategy {
		resp, err = a.call(url, method, body, headers...)
		if err == nil || !isRecoverable(err) {
			break
		}

		log.Printf("%s: http %s request error [%s], will retry in [%s]", logPrefix, method, err, retryWait)
		time.Sleep(retryWait)
	}

	return resp, err
}

// call - creates a new HTTP request and returns an HTTP response
func (a *Client) call(url string, method string, body io.Reader, headers ...map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Response{}, err
	}

	allHeaders := mergeHeaders(headers...)
	for key, value := range allHeaders {
		req.Header.Add(key, value)
	}

	resp, err := a.Client.Do(req)
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
