package fetch

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/code-gorilla-au/odize"
)

func TestClient_Patch_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	defaultHeaders := map[string]string{
		"default-header": "bar",
	}

	client := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := client.Patch("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Patch_no_retry_with_default_and_normal_headers(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	defaultHeaders := map[string]string{
		"default-header": "bar",
	}

	client := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := client.Patch("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestClient_Patch_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := client.Patch("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Delete_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := client.Delete("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Delete_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := client.Delete("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Delete_with_retry_with_default_and_normal_headers(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	defaultHeaders := map[string]string{
		"default-header": "bar",
	}

	client := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := client.Delete("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestClient_Put_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := client.Put("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Put_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := client.Put("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Put_with_retry_with_default_and_normal_headers(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	defaultHeaders := map[string]string{
		"default-header": "bar",
	}

	client := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := client.Put("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestClient_Get_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := client.Get("", headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}
func TestClient_Get_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := client.Get("", headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Get_no_retry_with_default_and_normal_headers(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	defaultHeaders := map[string]string{
		"default-header": "bar",
	}

	client := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := client.Get("", headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestClient_Post_with_retry_response_status_ok(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := client.Post("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Post_with_retry_response_should_try_once(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	_, _ = client.Post("", bytes.NewReader(nil), headers)
	odize.AssertEqual(t, m.Retries, 1)
}

func TestClient_Post_with_retry_response_should_try_twice(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusGatewayTimeout),
			StatusCode: http.StatusGatewayTimeout,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond, 1 * time.Nanosecond},
		Client:        &m,
	}

	_, _ = client.Post("", bytes.NewReader(nil), headers)
	odize.AssertEqual(t, m.Retries, 2)
}

func TestClient_Post_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := client.Post("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestClient_Post_empty_retry_list(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := &Client{
		RetryStrategy: []time.Duration{},
		Client:        &m,
	}

	_, err := client.Post("", bytes.NewReader(nil), headers)
	odize.AssertError(t, err)
}

func TestClient_Post_no_retry_with_default_and_normal_headers(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	defaultHeaders := map[string]string{
		"default-header": "bar",
	}

	client := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := client.Post("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestNew_with_default_retry(t *testing.T) {
	client := New(nil)
	odize.AssertEqual(t, client.RetryStrategy, setDefaultFetch().RetryStrategy)
}

func TestNew_with_default_header(t *testing.T) {
	client := New(nil)
	odize.AssertEqual(t, client.DefaultHeaders, setDefaultFetch().DefaultHeaders)
}

func TestNew_with_functional_options(t *testing.T) {
	expected := []time.Duration{1, 2}
	client := New(WithOpts(
		WithRetryStrategy(&expected),
	))
	odize.AssertEqual(t, client.RetryStrategy, expected)
}

func TestNew_with_options_headers(t *testing.T) {
	options := Options{
		WithRetry: false,
		DefaultHeaders: map[string]string{
			"foo": "bar",
		},
	}
	client := New(&options)
	odize.AssertEqual(t, client.DefaultHeaders, options.DefaultHeaders)
}

func TestNew_with_options_no_retry(t *testing.T) {
	options := Options{
		WithRetry: false,
		DefaultHeaders: map[string]string{
			"foo": "bar",
		},
	}
	client := New(&options)
	odize.AssertEqual(t, client.RetryStrategy, []time.Duration(nil))
}
func TestNew_with_options_with_retry(t *testing.T) {
	options := Options{
		WithRetry: true,
	}
	client := New(&options)
	odize.AssertEqual(t, client.RetryStrategy, setDefaultRetryStrategy())
}

func Test_mergeHeaders_should_merge_correctly(t *testing.T) {
	expected := map[string]string{
		"foo": "bar",
		"bin": "baz",
	}

	test := mergeHeaders(map[string]string{"foo": "bar"}, map[string]string{"bin": "baz"})

	odize.AssertEqual(t, expected, test)
}

func Test_mergeHeaders_empty_should_work(t *testing.T) {
	expected := map[string]string{}

	test := mergeHeaders()

	odize.AssertEqual(t, expected, test)
}
