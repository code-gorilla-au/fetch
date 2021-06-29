package fetch

import (
	"io"
	"net/http"
)

type MockHTTPClient struct {
	Resp    *http.Response
	Req     *http.Request
	Err     error
	ErrDo   bool
	ErrGet  bool
	ErrPost bool
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.Req = req
	if m.ErrDo {
		return nil, m.Err
	}
	if m.Resp == nil {
		return &http.Response{}, nil
	}
	return m.Resp, nil
}
func (m *MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	if m.ErrGet {
		return nil, m.Err
	}
	if m.Resp == nil {
		return &http.Response{}, nil
	}
	return m.Resp, nil
}
func (m *MockHTTPClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	if m.ErrPost {
		return nil, m.Err
	}
	if m.Resp == nil {
		return &http.Response{}, nil
	}
	return m.Resp, nil
}
