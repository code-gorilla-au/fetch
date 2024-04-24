package fetch

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/code-gorilla-au/odize"
)

func Test_call_POST_should_not_return_error_and_match_req(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	_, err := call(url, http.MethodPost, nil, &m, expectedHeaders)
	odize.AssertNoError(t, err)
	for key, value := range expectedHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	odize.AssertEqual(t, m.Req.URL.String(), url)
	odize.AssertEqual(t, m.Req.Method, http.MethodPost)

}

func Test_call_POST_4xx_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusBadRequest,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	var expectedErr *APIError

	_, err := call(url, http.MethodPost, nil, &m, expectedHeaders)
	if err == nil {
		t.Error("expected error, got none")
		return
	}
	odize.AssertTrue(t, errors.As(err, &expectedErr))
}

func Test_call_POST_5xx_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusInternalServerError,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	var expectedErr *APIError

	_, err := call(url, http.MethodPost, nil, &m, expectedHeaders)
	if err == nil {
		t.Error("expected error, got none")
		return
	}
	odize.AssertTrue(t, errors.As(err, &expectedErr))
}

func Test_call_GET_should_not_return_error_and_match_req(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	_, err := call(url, http.MethodGet, nil, &m, expectedHeaders)
	odize.AssertNoError(t, err)
	for key, value := range expectedHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	odize.AssertEqual(t, m.Req.URL.String(), url)
	odize.AssertEqual(t, m.Req.Method, http.MethodGet)

}
func Test_call_PUT_should_not_return_error_and_match_req(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	_, err := call(url, http.MethodPut, nil, &m, expectedHeaders)
	odize.AssertNoError(t, err)
	for key, value := range expectedHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	odize.AssertEqual(t, m.Req.URL.String(), url)
	odize.AssertEqual(t, m.Req.Method, http.MethodPut)

}

func Test_call_PATCH_should_not_return_error_and_match_req(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	_, err := call(url, http.MethodPatch, nil, &m, expectedHeaders)
	odize.AssertNoError(t, err)
	for key, value := range expectedHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	odize.AssertEqual(t, m.Req.URL.String(), url)
	odize.AssertEqual(t, m.Req.Method, http.MethodPatch)

}

func Test_call_DELETE_should_not_return_error_and_match_req(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
		},
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	_, err := call(url, http.MethodDelete, nil, &m, expectedHeaders)
	odize.AssertNoError(t, err)
	for key, value := range expectedHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	odize.AssertEqual(t, m.Req.URL.String(), url)
	odize.AssertEqual(t, m.Req.Method, http.MethodDelete)

}

func Test_call_body_should_match(t *testing.T) {
	m := MockHTTPClient{}

	body := map[string]string{
		"slap": "foo",
	}

	data, err := json.Marshal(&body)
	odize.AssertNoError(t, err)

	url := "foo"

	_, err = call(url, http.MethodPost, bytes.NewReader(data), &m, body)
	odize.AssertNoError(t, err)

	test := map[string]string{}
	err = json.NewDecoder(m.Req.Body).Decode(&test)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, test, body)

}

func Test_call_should_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		ErrDo: true,
		Err:   errors.New("expected error"),
	}

	expectedHeaders := map[string]string{
		"Auth": "/app/json",
	}

	url := "foo"

	_, err := call(url, http.MethodPost, nil, &m, expectedHeaders)
	odize.AssertTrue(t, errors.Is(err, m.Err))
}

func Test_callWithRetry_client_error_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		ErrDo: true,
		Err:   errors.New("expected error"),
	}

	_, err := callWithRetry("", http.MethodPost, nil, &m, []time.Duration{1 * time.Nanosecond})
	odize.AssertTrue(t, errors.Is(err, m.Err))
}

func Test_callWithRetry_4xx_client_error_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusBadRequest,
		},
	}
	var apiErr *APIError
	_, err := callWithRetry("", http.MethodPost, nil, &m, []time.Duration{1 * time.Nanosecond})
	odize.AssertTrue(t, errors.As(err, &apiErr))
	odize.AssertEqual(t, 1, m.Retries)
}

func Test_callWithRetry_5xx_client_error_retry_and_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusInternalServerError,
		},
	}
	var apiErr *APIError
	_, err := callWithRetry("", http.MethodPost, nil, &m, []time.Duration{1 * time.Nanosecond, 1 * time.Nanosecond})

	odize.AssertTrue(t, errors.As(err, &apiErr))
	odize.AssertEqual(t, 2, m.Retries)
}

func Test_callWithRetry_no_retries_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		ErrDo: true,
		Err:   errors.New("not expected error"),
	}

	_, err := callWithRetry("", http.MethodPost, nil, &m, []time.Duration{})
	odize.AssertTrue(t, errors.Is(err, ErrNoValidRetryStrategy))
}

func Test_callWithRetry_nill_retries_should_return_error(t *testing.T) {
	m := MockHTTPClient{
		ErrDo: true,
		Err:   errors.New("not expected error"),
	}

	_, err := callWithRetry("", http.MethodPost, nil, &m, nil)
	odize.AssertTrue(t, errors.Is(err, ErrNoValidRetryStrategy))
}

func Test_callWithRetry_should_return_response(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	resp, err := callWithRetry("", http.MethodGet, nil, &m, []time.Duration{1 * time.Nanosecond})
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	odize.AssertEqual(t, m.Req.Method, http.MethodGet)
}

func TestAxios_Patch_no_retry(t *testing.T) {
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

	axios := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := axios.Patch("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Patch_no_retry_with_default_and_normal_headers(t *testing.T) {
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

	axios := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := axios.Patch("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestAxios_Patch_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := axios.Patch("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Delete_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := axios.Delete("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Delete_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := axios.Delete("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Delete_with_retry_with_default_and_normal_headers(t *testing.T) {
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

	axios := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := axios.Delete("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestAxios_Put_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := axios.Put("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Put_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := axios.Put("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Put_with_retry_with_default_and_normal_headers(t *testing.T) {
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

	axios := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := axios.Put("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestAxios_Get_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := axios.Get("", headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}
func TestAxios_Get_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := axios.Get("", headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Get_no_retry_with_default_and_normal_headers(t *testing.T) {
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

	axios := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := axios.Get("", headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
	for key, value := range defaultHeaders {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
	for key, value := range headers {
		odize.AssertEqual(t, m.Req.Header.Get(key), value)
	}
}

func TestAxios_Post_with_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: []time.Duration{1 * time.Nanosecond},
		Client:        &m,
	}

	resp, err := axios.Post("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Post_no_retry(t *testing.T) {
	m := MockHTTPClient{
		Resp: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	axios := &Client{
		RetryStrategy: nil,
		Client:        &m,
	}

	resp, err := axios.Post("", bytes.NewReader(nil), headers)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, resp, m.Resp)
}

func TestAxios_Post_no_retry_with_default_and_normal_headers(t *testing.T) {
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

	axios := &Client{
		RetryStrategy:  nil,
		Client:         &m,
		DefaultHeaders: defaultHeaders,
	}

	resp, err := axios.Post("", bytes.NewReader(nil), headers)
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
	axios := New(nil)
	odize.AssertEqual(t, axios.RetryStrategy, setDefaultFetch().RetryStrategy)
}

func TestNew_with_default_header(t *testing.T) {
	axios := New(nil)
	odize.AssertEqual(t, axios.DefaultHeaders, setDefaultFetch().DefaultHeaders)
}

func TestNew_with_functional_options(t *testing.T) {
	expected := []time.Duration{1, 2}
	axios := New(WithOpts(
		WithRetryStrategy(&expected),
	))
	odize.AssertEqual(t, axios.RetryStrategy, expected)
}

func TestNew_with_options_headers(t *testing.T) {
	options := Options{
		WithRetry: false,
		DefaultHeaders: map[string]string{
			"foo": "bar",
		},
	}
	axios := New(&options)
	odize.AssertEqual(t, axios.DefaultHeaders, options.DefaultHeaders)
}

func TestNew_with_options_no_retry(t *testing.T) {
	options := Options{
		WithRetry: false,
		DefaultHeaders: map[string]string{
			"foo": "bar",
		},
	}
	axios := New(&options)
	odize.AssertEqual(t, axios.RetryStrategy, []time.Duration(nil))
}
func TestNew_with_options_with_retry(t *testing.T) {
	options := Options{
		WithRetry: true,
	}
	axios := New(&options)
	odize.AssertEqual(t, axios.RetryStrategy, setDefaultRetryStrategy())
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
