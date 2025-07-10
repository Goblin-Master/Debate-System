package handler

import (
	"Debate-System/user/internal/logic/user"
	"Debate-System/user/internal/response" // ①
	"Debate-System/user/internal/svc"
	"Debate-System/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err) //②
	}
}
