package ws

import (
	"Debate-System/chat/internal/config"
	"Debate-System/chat/internal/svc"
	"Debate-System/utils/jwtx"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strconv"
)

func Server(c config.Config, svcCtx *svc.ServiceContext) {
	//// ✅ 正确方式：启动链路追踪
	//trace.StartAgent(trace.Config{})
	//defer trace.StopAgent()

	hub := NewHub(svcCtx)

	server := rest.MustNewServer(c.WSServer)
	defer server.Stop()

	// 关键：注册 /ws 路由，完全沿用你的 handler 逻辑
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			logger := logx.WithContext(r.Context())
			user_id, err := jwtx.GetUserID(r.Context())
			if err != nil {
				logger.Error(err)
				w.Write(USER_STATE_ERROR)
				return
			}
			conversation_id := r.URL.Query().Get(COVERSATION_ID)
			if conversation_id == "" {
				logger.Error(CONVERSATION_ID_EMPTY)
				w.Write(CONVERSATION_NOT_EXIST)
				return
			}
			conn, err := (&websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool { return true },
			}).Upgrade(w, r, nil)
			if err != nil {
				logger.Error(err)
				w.Write(WS_INIT_ERROR)
				return
			}
			hub.AddConn(r.Context(), strconv.FormatInt(user_id, 10), conversation_id, conn)
		},
	},
		rest.WithJwt(svcCtx.Config.Auth.AccessSecret),
	)

	fmt.Printf("Starting ws server at %s:%d...\n", c.WSServer.Host, c.WSServer.Port)
	server.Start()
}
