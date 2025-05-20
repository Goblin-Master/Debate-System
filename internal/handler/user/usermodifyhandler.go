package handler

import (
	"Debate-System/internal/logic/user"
	"Debate-System/internal/response" // ①
	"Debate-System/internal/svc"
	"Debate-System/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserModifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserModifyLogic(r.Context(), svcCtx)
		resp, err := l.UserModify(&req)
		response.Response(w, resp, err) //②

	}
}
