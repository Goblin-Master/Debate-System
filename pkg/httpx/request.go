package httpx

import (
	"Debate-System/pkg/iox"
	"context"
	"io"
	"net/http"
)

type Request struct {
	req    *http.Request
	err    error
	client *http.Client
}

func NewRequest(ctx context.Context, method, url string) *Request {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	return &Request{
		req:    req,
		err:    err,
		client: http.DefaultClient,
	}
}

// JSONBody 使用 JSON body
func (req *Request) JSONBody(val any) *Request {
	if req.err != nil {
		return req
	}
	req.req.Body = io.NopCloser(iox.NewJSONReader(val))
	req.req.Header.Set("Content-Type", "application/json")
	return req
}

func (req *Request) Client(cli *http.Client) *Request {
	req.client = cli
	return req
}

func (req *Request) AddHeader(key string, value string) *Request {
	if req.err != nil {
		return req
	}
	req.req.Header.Add(key, value)
	return req
}

// AddParam 添加查询参数
// 这个方法性能不好，但是好用
func (req *Request) AddParam(key string, value string) *Request {
	if req.err != nil {
		return req
	}
	q := req.req.URL.Query()
	q.Add(key, value)
	req.req.URL.RawQuery = q.Encode()
	return req
}

// URL 返回已经拼接好查询参数的完整 URL
func (req *Request) URL() string {
	if req.err != nil || req.req == nil {
		return ""
	}
	return req.req.URL.String()
}
func (req *Request) Do() *Response {
	if req.err != nil {
		return &Response{
			err: req.err,
		}
	}
	resp, err := req.client.Do(req.req)
	return &Response{
		Response: resp,
		err:      err,
	}
}
