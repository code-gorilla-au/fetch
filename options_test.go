package fetch

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithRetryStrategy_with_nil_strategy(t *testing.T) {
	fn := WithRetryStrategy(nil)

	options := Options{}
	err := fn(&options)
	assert.NoError(t, err)
	assert.Nil(t, options.RetryStrategy)
}

func TestWithRetryStrategy_with_custom_strategy(t *testing.T) {
	st := []time.Duration{1, 2}
	fn := WithRetryStrategy(&st)

	options := Options{}
	err := fn(&options)
	assert.NoError(t, err)
	assert.Equal(t, &st, options.RetryStrategy)
}

func TestWithHTTPClient_with_nil_client(t *testing.T) {
	fn := WithHTTPClient(nil)

	options := Options{}
	err := fn(&options)
	assert.NoError(t, err)
	assert.Nil(t, options.HTTPClient)
}

func TestWithHTTPClient_with_custom_client(t *testing.T) {
	fn := WithHTTPClient(&http.Client{})

	options := Options{}
	err := fn(&options)
	assert.NoError(t, err)
	assert.Equal(t, &http.Client{}, options.HTTPClient)
}

func TestWithOpts_no_options(t *testing.T) {
	options := WithOpts(nil)
	assert.Nil(t, options.HTTPClient)
	assert.Nil(t, options.RetryStrategy)
}

func TestWithOpts_with_custom_retry(t *testing.T) {
	st := []time.Duration{1, 2}
	options := WithOpts(WithRetryStrategy(&st))
	assert.Equal(t, &st, options.RetryStrategy)
	assert.Nil(t, options.HTTPClient)
}

func TestWithOpts_with_custom_client(t *testing.T) {
	cl := http.Client{}
	options := WithOpts(WithHTTPClient(&cl))
	assert.Equal(t, &cl, options.HTTPClient)
	assert.Nil(t, options.RetryStrategy)
}
