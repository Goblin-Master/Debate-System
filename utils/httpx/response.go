package httpx

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	*http.Response
	err error
}

// JSONScan 将 Body 按照 JSON 反序列化为结构体
func (r *Response) JSONScan(val any) error {
	if r.err != nil {
		return r.err
	}
	err := json.NewDecoder(r.Body).Decode(val)
	return err
}
