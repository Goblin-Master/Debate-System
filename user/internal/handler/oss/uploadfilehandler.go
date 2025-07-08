package handler

import (
	"Debate-System/user/internal/logic/oss"
	"Debate-System/user/internal/response"
	"Debate-System/user/internal/svc"
	"errors"
	"net/http"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if svcCtx.Config.OSS.Enable == false {
			response.Response(w, nil, errors.New("OSS未启用"))
			return
		}
		_, file, err := r.FormFile("file")
		if err != nil {
			response.Response(w, nil, err)
			return
		}
		l := oss.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(file)
		response.Response(w, resp, err) //②

	}
}
