package handler

import (
	"Debate-System/response" // ①
	"Debate-System/reward/internal/logic/score"
	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ModifyScoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyScoreReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := score.NewModifyScoreLogic(r.Context(), svcCtx)
		resp, err := l.ModifyScore(&req)
		response.Response(w, resp, err) //②

	}
}
