package ai

import (
	"Debate-System/chat/internal/svc"
	"Debate-System/chat/internal/types"
	"Debate-System/pkg/httpx"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type CozeCreateConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCozeCreateConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CozeCreateConversationLogic {
	return &CozeCreateConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var (
	CONVERSATION_CREATE_FAILED = errors.New("创建辩论失败")
)

func (l *CozeCreateConversationLogic) CozeCreateConversation(req *types.CozeCreateConversationReq) (resp *types.CozeCreateConversationResp, err error) {
	request := httpx.NewRequest(l.ctx, "POST", "https://api.coze.cn/v1/conversation/create")
	request.AddHeader("Content-Type", "application/json")
	request.AddHeader("Authorization", fmt.Sprintf("Bearer %s", l.svcCtx.Config.Coze.Token))
	type reqBody struct {
		BotID    string `json:"bot_id"`
		Messages []struct {
			Content     string `json:"content"`
			ContentType string `json:"content_type"`
			Type        string `json:"type"`
			Role        string `json:"role"`
		} `json:"messages"`
	}
	request.JSONBody(reqBody{
		BotID: l.svcCtx.Config.Coze.BotID,
		Messages: []struct {
			Content     string `json:"content"`
			ContentType string `json:"content_type"`
			Type        string `json:"type"`
			Role        string `json:"role"`
		}{
			{
				Content:     fmt.Sprintf("这是我们接下来要进行的辩论的题目:%s", req.Theme),
				ContentType: "text",
				Type:        "text",
				Role:        "assistant",
			},
		},
	})
	client := &http.Client{}
	request.Client(client)
	response := request.Do()
	if response.StatusCode != http.StatusOK {
		l.Logger.Error(response.StringBody())
		return nil, CONVERSATION_CREATE_FAILED
	}
	type respBody struct {
		Data struct {
			ConversationID string `json:"id"`
			Ctime          int64  `json:"created_at"`
		} `json:"data"`
	}
	var body respBody
	if err := response.JSONScan(&body); err != nil {
		l.Logger.Error(err)
		return nil, CONVERSATION_CREATE_FAILED
	}
	resp = &types.CozeCreateConversationResp{
		ConversationID: body.Data.ConversationID,
		Ctime:          body.Data.Ctime,
	}
	return resp, nil
}
