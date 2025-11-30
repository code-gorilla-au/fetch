// Package fetch provides a simple fetch client with a built-in backup / retry strategy.
// Provides a interface for cancelable HTTP requests.
package fetch

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

// New initialises and returns a new Client instance with the provided options or default configurations if nil.
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

// Get sends an HTTP GET request to the specified URL with optional headers and returns the HTTP response or an error.
//
// Example:
//
//	var apiErr *fetch.APIError
//
//	resp, err := client.Get(url, nil)
//	if err != nil {
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	ctx := context.Background()
	return a.do(ctx, url, http.MethodGet, nil, headers)
}

// Post sends an HTTP POST request to the specified URL with a body and optional headers, returning the HTTP response or an error.
//
// Example:
//
//	var apiErr *fetch.APIError
//
//	resp, err := client.Post(url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	ctx := context.Background()
	return a.do(ctx, url, http.MethodPost, body, headers)
}

// Put sends an HTTP PUT request to the specified URL with a body and optional headers, returning the HTTP response or an error.
//
// Example:
//
//	var apiErr *fetch.APIError
//
//	resp, err := client.Put(url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) Put(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	ctx := context.Background()
	return a.do(ctx, url, http.MethodPut, body, headers)
}

// Delete sends an HTTP DELETE request to the specified URL with a body and optional headers, returning the response or an error.
//
// Example:
//
//	var apiErr *fetch.APIError
//
//	resp, err := client.Delete(url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) Delete(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	ctx := context.Background()
	return a.do(ctx, url, http.MethodDelete, body, headers)
}

// Patch sends an HTTP PATCH request to the specified URL with a body and optional headers, returning the response or an error.
// Example:
//
//	var apiErr *fetch.APIError
//
//	resp, err := client.Patch(url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) Patch(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	ctx := context.Background()
	return a.do(ctx, url, http.MethodPatch, body, headers)
}

// GetCtx sends a cancelable HTTP GET request to the specified URL with context and optional headers, returning the response or an error.
//
// Example:
//
//	var apiErr *fetch.APIError
//	ctx, cancel := context.WithCancel(context.Background())
//
//	resp, err := client.GetCtx(ctx,url, nil)
//	if err != nil {
//		if errors.is(err, context.Canceled) {
//			// Handle context cancelled
//		}
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) GetCtx(ctx context.Context, url string, headers map[string]string) (*http.Response, error) {
	return a.do(ctx, url, http.MethodGet, nil, headers)
}

// PostCtx sends a cancelable HTTP POST request to the specified URL with context, body, and headers, returning a response or error.
//
// Example:
//
//	var apiErr *fetch.APIError
//	ctx, cancel := context.WithCancel(context.Background())
//
//	resp, err := client.PostCtx(ctx,url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.is(err, context.Canceled) {
//			// Handle context cancelled
//		}
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) PostCtx(ctx context.Context, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(ctx, url, http.MethodPost, body, headers)
}

// PutCtx sends a cancelable HTTP PUT request to the specified URL with context, body, and headers, returning a response or error.
//
// Example:
//
//	var apiErr *fetch.APIError
//	ctx, cancel := context.WithCancel(context.Background())
//
//	resp, err := client.PutCtx(ctx,url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.is(err, context.Canceled) {
//			// Handle context cancelled
//		}
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) PutCtx(ctx context.Context, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(ctx, url, http.MethodPut, body, headers)
}

// DeleteCtx sends a cancelable HTTP DELETE request to the specified URL with context, body, and headers, returning a response or error.
//
// Example:
//
//	var apiErr *fetch.APIError
//	ctx, cancel := context.WithCancel(context.Background())
//
//	resp, err := client.DeleteCtx(ctx,url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.is(err, context.Canceled) {
//			// Handle context cancelled
//		}
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) DeleteCtx(ctx context.Context, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(ctx, url, http.MethodDelete, body, headers)
}

// PatchCtx sends an HTTP PATCH request to the specified URL with the provided context, body, and headers.
//
// Example:
//
//	var apiErr *fetch.APIError
//	ctx, cancel := context.WithCancel(context.Background())
//
//	resp, err := client.PatchCtx(ctx,url, bytes.NewReader([]byte(`{"hello": "world"}`)), nil)
//	if err != nil {
//		if errors.is(err, context.Canceled) {
//			// Handle context cancelled
//		}
//		if errors.As(err, &apiErr) {
//			fmt.Println("API Response error", apiErr)
//		}
//		// Handle non-API Error
//	}
func (a *Client) PatchCtx(ctx context.Context, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	return a.do(ctx, url, http.MethodPatch, body, headers)
}

// do - make http call with the provided configuration
func (a *Client) do(ctx context.Context, url string, method string, body io.Reader, headers map[string]string) (*http.Response, error) {
	if a.RetryStrategy == nil {
		return a.call(ctx, url, method, body, headers, a.DefaultHeaders)
	}

	return a.callWithRetry(ctx, url, method, body, headers, a.DefaultHeaders)
}

// callWithRetry - wrap the call method with the retry strategy
func (a *Client) callWithRetry(ctx context.Context, url string, method string, body io.Reader, headers ...map[string]string) (*http.Response, error) {
	logPrefix := "fetch: callWithRetry"
	var resp *http.Response
	var err error

	if len(a.RetryStrategy) == 0 {
		return resp, ErrNoValidRetryStrategy
	}

	for _, retryWait := range a.RetryStrategy {
		resp, err = a.call(ctx, url, method, body, headers...)

		if err == nil || !isRecoverable(err) {
			if errors.Is(err, context.Canceled) {
				log.Printf("%s: http %s request canceled", logPrefix, method)
			}

			break
		}

		log.Printf("%s: http %s request error [%s], will retry in [%s]", logPrefix, method, err, retryWait)
		time.Sleep(retryWait)
	}

	return resp, err
}

// call - creates a new HTTP request and returns an HTTP response
func (a *Client) call(ctx context.Context, url string, method string, body io.Reader, headers ...map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return &http.Response{}, err
	}

	allHeaders := mergeHeaders(headers...)
	for key, value := range allHeaders {
		req.Header.Add(key, value)
	}

	log.Println("request", req == nil)

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
