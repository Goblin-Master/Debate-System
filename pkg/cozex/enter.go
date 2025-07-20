package cozex

import (
	"github.com/coze-dev/coze-go"
	"net/http"
	"time"
)

type Coze struct {
	BotID string
	Token string
}

func InitCozeClient(c Coze) coze.CozeAPI {
	token := c.Token
	authClient := coze.NewTokenAuth(token)

	customClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	client := coze.NewCozeAPI(authClient,
		coze.WithBaseURL("https://api.coze.cn"),
		coze.WithHttpClient(customClient),
	)
	return client
}
