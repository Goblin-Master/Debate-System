package httpx

import (
	"context"
	"errors"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequest_Client(t *testing.T) {
	req := NewRequest(context.Background(), http.MethodPost, "/abc")
	assert.Equal(t, http.DefaultClient, req.client)
	cli := &http.Client{}
	req = req.Client(&http.Client{})
	assert.Equal(t, cli, req.client)
}

func TestRequest_JSONBody(t *testing.T) {
	req := NewRequest(context.Background(), http.MethodPost, "/abc")
	assert.Nil(t, req.req.Body)
	req = req.JSONBody(User{})
	assert.NotNil(t, req.req.Body)
	assert.Equal(t, "application/json", req.req.Header.Get("Content-Type"))

	req2 := NewRequest(context.Background(), http.MethodGet, "://localhost:80/a")
	assert.NotNil(t, req2.err)
	assert.Nil(t, req2.req)
}

func TestRequest_Do(t *testing.T) {
	l, err := net.Listen("unix", "/tmp/test.sock")
	require.NoError(t, err)
	server := http.Server{}
	go func() {
		http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = writer.Write([]byte("OK"))
		})
		_ = server.Serve(l)
	}()
	defer func() {
		_ = l.Close()
	}()
	testCases := []struct {
		name    string
		req     func() *Request
		wantErr error
	}{
		{
			name: "构造请求的时候有 error",
			req: func() *Request {
				return &Request{
					err: errors.New("mock error"),
				}
			},
			wantErr: errors.New("mock error"),
		},
		{
			name: "成功",
			req: func() *Request {
				req := NewRequest(context.Background(), http.MethodGet, "http://localhost:8081/hello")
				return req.Client(&http.Client{
					Transport: &http.Transport{
						DialContext: func(ctx context.Context,
							network, addr string) (net.Conn, error) {
							return net.Dial("unix", "/tmp/test.sock")
						},
					},
				})
			},
		},
	}

	// 确保前面的 http 端口启动成功
	time.Sleep(time.Second)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.req()
			resp := req.Do()
			assert.Equal(t, tc.wantErr, resp.err)
		})
	}
}

func TestRequest_AddParam(t *testing.T) {
	req := NewRequest(context.Background(),
		http.MethodGet, "http://localhost").
		AddParam("key1", "value1").
		AddParam("key2", "value2")
	assert.Equal(t, "http://localhost?key1=value1&key2=value2", req.req.URL.String())

	req2 := NewRequest(context.Background(), http.MethodGet, "://localhost:80/a")
	assert.NotNil(t, req2.err)
	assert.Nil(t, req2.req)
}

func TestRequestAddHeader(t *testing.T) {
	req := NewRequest(context.Background(),
		http.MethodGet, "http://localhost").
		AddHeader("head1", "val1").AddHeader("head1", "val2")
	vals := req.req.Header.Values("head1")
	assert.Equal(t, []string{"val1", "val2"}, vals)

	req2 := NewRequest(context.Background(), http.MethodGet, "://localhost:80/a")
	assert.NotNil(t, req2.err)
	assert.Nil(t, req2.req)
}

type User struct {
	Name string
}
