package handler

import (
	"Debate-System/internal/logic/oss"
	"Debate-System/internal/response" // ①
	"Debate-System/internal/svc"
	"net/http"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
