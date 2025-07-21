package ws

import (
	"Debate-System/chat/internal/svc"
	"Debate-System/utils/syncx"
	"context"
	"github.com/coze-dev/coze-go"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Hub struct {
	// syncx.Map 是我对 sync.Map 的一个简单封装
	conns  *syncx.Map[string, *websocket.Conn]
	svcCtx *svc.ServiceContext
}

func NewHub(svcCtx *svc.ServiceContext) *Hub {
	return &Hub{
		svcCtx: svcCtx,
		conns:  &syncx.Map[string, *websocket.Conn]{},
	}
}

var (
	USER_STATE_ERROR       = []byte("用户信息异常，请重新登陆")
	WS_INIT_ERROR          = []byte("初始化 websocket 失败")
	CONVERSATION_NOT_EXIST = []byte("该辩论不存在")
	CONVERSATION_ID_EMPTY  = "conversation_id 为空"
	COVERSATION_ID         = "conversation_id"
)

func (h *Hub) AddConn(ctx context.Context, user_id, conversation_id string, c *websocket.Conn) {
	h.conns.Store(user_id, c)
	// ✅ 用带 trace 的 ctx 重新生成 logger，日志就带 trace-id
	logger := logx.WithContext(ctx)
	go func() {
		for {
			typ, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNoStatusReceived) {
					//客户端断开连接
					h.conns.Delete(user_id)
				} else {
					logger.Error(err)
				}
				return
			}
			// 开始转发
			h.conns.Range(func(key string, value *websocket.Conn) bool {
				if key == user_id {
					//h.cozeChat(conversation_id, user_id, string(message))
					logger.Info("用户开始对话", user_id, conversation_id, string(message))
					err = value.WriteMessage(typ, message)
					if err != nil {
						logger.Error(err)
					}
				}
				// 返回 true，确保会继续往后遍历
				return true
			})
		}
	}()
}

//func (h *Hub) Server() {
//	//允许跨域连接，CheckOrigin 返回true
//	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
//	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
//		// 没有额外的 header 的设置
//		user_id, err := jwtx.ParseToken(h.svcCtx.Config.Auth.AccessSecret, request)
//		if err != nil {
//			h.Logger.Error(err)
//			writer.Write(USER_STATE_ERROR)
//			return
//		}
//		conversation_id := request.URL.Query().Get(COVERSATION_ID)
//		if conversation_id == "" {
//			h.Logger.Error(CONVERSATION_ID_EMPTY)
//			writer.Write(CONVERSATION_NOT_EXIST)
//			return
//		}
//		conn, err := upgrader.Upgrade(writer, request, nil)
//		if err != nil {
//			h.Logger.Error(err)
//			writer.Write(WS_INIT_ERROR)
//			return
//		}
//		h.Logger.Info("用户加入成功", user_id, conversation_id)
//		h.AddConn(strconv.FormatInt(user_id, 10), conversation_id, conn)
//	})
//	http.ListenAndServe(h.svcCtx.Config.WS.Addr, nil)
//}

func (h *Hub) cozeChat(conversation_id, user_id, content string) (coze.Stream[coze.ChatEvent], error) {
	request := &coze.CreateChatsReq{
		ConversationID: conversation_id,
		BotID:          h.svcCtx.Config.Coze.BotID,
		UserID:         user_id,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText(content, nil),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return h.svcCtx.CozeClient.Chat.Stream(ctx, request)
}
