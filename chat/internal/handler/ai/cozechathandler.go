package handler

import (
	"Debate-System/chat/internal/logic/ai"
	"Debate-System/chat/internal/svc"
	"Debate-System/chat/internal/types"
	"Debate-System/response" // ①
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CozeChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CozeChatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := ai.NewCozeChatLogic(r.Context(), svcCtx)
		resp, err := l.CozeChat(&req)
		response.Response(w, resp, err) //②

	}
}
