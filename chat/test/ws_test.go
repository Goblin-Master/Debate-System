package test

import (
	"Debate-System/chat/internal/config"
	"Debate-System/chat/internal/svc"
	"Debate-System/chat/ws"
	"Debate-System/utils/jwtx"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strconv"
	"testing"
)

var configFile = flag.String("f", "J:\\Debate-System\\chat\\etc\\chat-api.yaml", "the config file")

func Test_WS_Server(t *testing.T) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()

	//// ✅ 正确方式：启动链路追踪
	//trace.StartAgent(trace.Config{})
	//defer trace.StopAgent()

	ctx := svc.NewServiceContext(c)
	hub := ws.NewHub(ctx)

	server := rest.MustNewServer(c.HttpServer)
	defer server.Stop()

	// 关键：注册 /ws 路由，完全沿用你的 handler 逻辑
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			user_id, err := jwtx.GetUserID(r.Context())
			if err != nil {
				w.Write(ws.USER_STATE_ERROR)
				return
			}
			conversation_id := r.URL.Query().Get(ws.COVERSATION_ID)
			if conversation_id == "" {
				w.Write(ws.CONVERSATION_NOT_EXIST)
				return
			}
			conn, err := (&websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool { return true },
			}).Upgrade(w, r, nil)
			if err != nil {
				w.Write(ws.WS_INIT_ERROR)
				return
			}
			hub.AddConn(r.Context(), strconv.FormatInt(user_id, 10), conversation_id, conn)
		},
	},
		rest.WithJwt(ctx.Config.Auth.AccessSecret),
	)
	server.Start()
}
