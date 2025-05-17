package test

import (
	"Debate-System/utils/jwtx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwt(t *testing.T) {
	token, err := jwtx.GenToken(jwtx.Claims{
		Auth:   jwtx.Auth{AccessExpire: 864000000, AccessSecret: "Debate-System"},
		UserID: 1923059961114398720,
	})
	assert.NoError(t, err)
	t.Log(token)
}
