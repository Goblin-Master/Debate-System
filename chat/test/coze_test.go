package test

import (
	"Debate-System/pkg/httpx"
	"context"
	"errors"
	"fmt"
	"github.com/coze-dev/coze-go"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

// TestGetCode
//
//	@Description: 用于生成获取code的url，code只能用一次，且只用5分钟的有效期
//	@param t
func TestGetCode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	req := httpx.NewRequest(ctx, "GET", "https://www.coze.cn/api/permission/oauth2/authorize")
	req.AddParam("response_type", "code")
	req.AddParam("client_id", "")
	req.AddParam("redirect_uri", "http://localhost:8080")
	req.AddParam("state", "test")
	fmt.Println("获取code的url：", req.URL())
}

// TestCetToken
//
//	@Description: 获取token信息，atoken只有15分钟有效期，rtoken只有30天有效期
//	@param t
func TestCetToken(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	req := httpx.NewRequest(ctx, "POST", "https://api.coze.cn/api/permission/oauth2/token")
	req.AddHeader("Authorization", "Bearer client_secret")
	req.AddHeader("Content-Type", "application/json")
	req.JSONBody(map[string]string{
		"grant_type":   "authorization_code",
		"code":         "code_qQ6bw0hig5dOlWiwfGK5hlJ8EDgizrN1LwD9BskHNND49aXx",
		"redirect_uri": "http://localhost:8080",
		"client_id":    "",
	})
	client := &http.Client{}
	req.Client(client)
	resp := req.Do()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	type Response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	var data Response
	err := resp.JSONScan(&data)
	assert.NoError(t, err)
	fmt.Println(data)
}

// TestRefreshToken
//
//	@Description: 用rtoken刷新atoken，刷新成功后，用新的atoken访问api，同时旧的rtoken也没有用了
//	@param t
func TestRefreshToken(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	req := httpx.NewRequest(ctx, "POST", "https://api.coze.cn/api/permission/oauth2/token")
	req.AddHeader("Authorization", "Bearer client_secret")
	req.AddHeader("Content-Type", "application/json")
	req.JSONBody(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": "",
		"client_id":     "",
	})
	client := &http.Client{}
	req.Client(client)
	resp := req.Do()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	type Response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	var data Response
	err := resp.JSONScan(&data)
	assert.NoError(t, err)
	fmt.Println(data)
}

func TestChat(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get an access_token through personal access token or oauth.
	token := ""
	botID := "7528018417983504425"
	userID := "123"

	authCli := coze.NewTokenAuth(token)

	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"), coze.WithHttpClient(&http.Client{
		Timeout: time.Minute * 2,
	}))

	//
	// Step one, create chats
	// Call the coze.chats().stream() method to create a chats. The create method is a streaming
	// chats and will return a Flowable ChatEvent. Developers should iterate the iterator to get
	// chats event and handle them.
	// //
	req := &coze.CreateChatsReq{
		BotID:  botID,
		UserID: userID,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText("What can you do?", nil),
		},
	}

	resp, err := cozeCli.Chat.Stream(ctx, req)
	if err != nil {
		fmt.Printf("Error starting chats: %v\n", err)
		return
	}

	defer resp.Close()
	for {
		event, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if event.Event == coze.ChatEventConversationMessageDelta {
			fmt.Print(event.Message.Content, "#")
		} else if event.Event == coze.ChatEventConversationChatCompleted {
			fmt.Printf("Token usage:%d\n", event.Chat.Usage.TokenCount)
		} else {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("done, log:%s\n", resp.Response().LogID())
}
