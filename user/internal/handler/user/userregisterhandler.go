package handler

import (
	"Debate-System/response"
	"Debate-System/user/internal/logic/user"
	"Debate-System/user/internal/svc"
	"Debate-System/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		response.Response(w, resp, err) //②

	}
}
