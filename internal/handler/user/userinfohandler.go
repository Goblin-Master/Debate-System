package handler

import (
	"Debate-System/internal/logic/user"
	"Debate-System/internal/response" // ①
	"Debate-System/internal/svc"
	"net/http"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err) //②

	}
}
