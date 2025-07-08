package iox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONReader(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		val     any
		wantRes []byte
		wantN   int
		wantErr error
	}{
		{
			name:    "正常读取",
			input:   make([]byte, 10),
			val:     User{Name: "Tom"},
			wantN:   10,
			wantRes: []byte(`{"name":"T`), // 修正了 JSON 序列化结果
		},
		{
			name:    "输入 nil",
			input:   make([]byte, 7),
			val:     nil,
			wantN:   4,
			wantRes: []byte("null"), // 修正了 nil 的序列化结果
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := NewJSONReader(tc.val)
			n, err := reader.Read(tc.input)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantN, n)
			assert.Equal(t, tc.wantRes, tc.input[:n]) // 只比较实际读取的字节
		})
	}
}

type User struct {
	Name string `json:"name"`
}
