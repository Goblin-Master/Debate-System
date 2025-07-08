package httpx

import (
	"Debate-System/utils/iox"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_JSONScan(t *testing.T) {
	testCases := []struct {
		name    string
		resp    *Response
		wantVal User
		wantErr error
	}{
		{
			name: "scan成功",
			resp: &Response{
				Response: &http.Response{
					Body: io.NopCloser(iox.NewJSONReader(User{Name: "Tom"})),
				},
			},
			wantVal: User{
				Name: "Tom",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var u User
			err := tc.resp.JSONScan(&u)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, u)
		})
	}
}
