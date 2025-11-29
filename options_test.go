package fetch

import (
	"net/http"
	"testing"
	"time"

	"github.com/code-gorilla-au/odize"
)

func TestWithRetryStrategy_with_nil_strategy(t *testing.T) {
	fn := WithRetryStrategy(nil)

	options := Options{}
	err := fn(&options)
	odize.AssertNoError(t, err)
	odize.AssertNil(t, options.RetryStrategy)
}

func TestWithRetryStrategy_with_custom_strategy(t *testing.T) {
	st := []time.Duration{1, 2}
	fn := WithRetryStrategy(&st)

	options := Options{}
	err := fn(&options)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, &st, options.RetryStrategy)
	odize.AssertTrue(t, options.WithRetry)
}

func TestWithHTTPClient_with_nil_client(t *testing.T) {
	fn := WithHTTPClient(nil)

	options := Options{}
	err := fn(&options)
	odize.AssertNoError(t, err)
	odize.AssertNil(t, options.HTTPClient)
}

func TestWithHTTPClient_with_custom_client(t *testing.T) {
	fn := WithHTTPClient(&http.Client{})

	options := Options{}
	err := fn(&options)
	odize.AssertNoError(t, err)
	odize.AssertEqual(t, &http.Client{}, options.HTTPClient)
}

func TestWithOpts_no_options(t *testing.T) {
	options := WithOpts(nil)
	odize.AssertNil(t, options.HTTPClient)
	odize.AssertNil(t, options.RetryStrategy)
}

func TestWithOpts_with_custom_retry(t *testing.T) {
	st := []time.Duration{1, 2}
	options := WithOpts(WithRetryStrategy(&[]time.Duration{1, 2}))
	odize.AssertEqual(t, &st, options.RetryStrategy)
	odize.AssertNil(t, options.HTTPClient)
}

func TestWithOpts_with_custom_client(t *testing.T) {
	cl := http.Client{}
	options := WithOpts(WithHTTPClient(&cl))
	odize.AssertEqual(t, &cl, options.HTTPClient)
	odize.AssertNil(t, options.RetryStrategy)
}

func TestWithOpts_with_custom_retry_and_client(t *testing.T) {
	st := []time.Duration{1, 2}
	cl := http.Client{}
	options := WithOpts(WithHTTPClient(&cl), WithRetryStrategy(&[]time.Duration{1, 2}))
	odize.AssertEqual(t, &st, options.RetryStrategy)
}

func TestWithOpts_with_multiple_options(t *testing.T) {
	st := []time.Duration{1, 2}
	cl := http.Client{}
	options := WithOpts(WithHTTPClient(&cl), WithRetryStrategy(&[]time.Duration{1, 2}))
	odize.AssertEqual(t, &st, options.RetryStrategy)
}

func TestWithOpts_with_nil_options(t *testing.T) {
	options := WithOpts(nil)
	odize.AssertNil(t, options.HTTPClient)
	odize.AssertNil(t, options.RetryStrategy)
}

func TestWithOpts_with_empty_options(t *testing.T) {
	options := WithOpts()
	odize.AssertNil(t, options.HTTPClient)
}

func TestWithOpts_with_default_retry_strategy(t *testing.T) {
	options := WithOpts(WithDefaultRetryStrategy())
	odize.AssertTrue(t, options.WithRetry)
}

func TestWithOpts_with_headers(t *testing.T) {
	headers := map[string]string{
		"foo": "bar",
	}
	options := WithOpts(WithHeaders(headers))
	odize.AssertEqual(t, headers, options.DefaultHeaders)
}
