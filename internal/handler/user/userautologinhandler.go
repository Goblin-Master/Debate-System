package handler

import (
	"Debate-System/internal/logic/user"
	"Debate-System/internal/response" // ①
	"Debate-System/internal/svc"
	"net/http"
)

func UserAutoLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUserAutoLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserAutoLogin()
		response.Response(w, resp, err) //②

	}
}
