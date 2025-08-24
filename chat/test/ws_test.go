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

func TestMain(m *testing.M) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()

	ctx := svc.NewServiceContext(c)
	hub := ws.NewHub(ctx)

	server := rest.MustNewServer(c.WSServer)
	defer server.Stop()

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

	go server.Start() // 用 goroutine 启动
	select {}         // 阻塞住，防止退出
}
