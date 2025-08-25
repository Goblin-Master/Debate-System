package ws

import (
	"Debate-System/chat/internal/svc"
	"Debate-System/utils/syncx"
	"context"
	"errors"
	"fmt"
	"github.com/coze-dev/coze-go"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
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

// coze流式响应自己包装的code
const (
	FAIL int = iota - 1
	SUCCESS
	END
)

func (h *Hub) AddConn(ctx context.Context, user_id, conversation_id string, c *websocket.Conn) {
	h.conns.Store(user_id, c)
	// ✅ 用带 trace 的 ctx 重新生成 logger，日志就带 trace-id
	logger := logx.WithContext(ctx)

	go func() {
		for {
			typ, message, err := c.ReadMessage()
			//监视是否在线
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
				// true就是继续下一个，false直接终止全部,由于辩论只需要自己知道，所以异常直接return false
				if key == user_id {
					//为每次对话加一个超时控制时间
					msgCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
					defer cancel()
					//coze对话
					h.cozeChat(msgCtx, logger, value, typ, conversation_id, user_id, string(message))
				}
				// 返回 true，确保会继续往后遍历
				return true
			})
		}
	}()
}

func (h *Hub) cozeChat(ctx context.Context, l logx.Logger, c *websocket.Conn, typ int, conversation_id, user_id, content string) {
	request := &coze.CreateChatsReq{
		ConversationID: conversation_id,
		BotID:          h.svcCtx.Config.Coze.BotID,
		UserID:         user_id,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText(content, nil),
		},
	}
	resp, err := h.svcCtx.CozeClient.Chat.Stream(ctx, request)
	if err != nil {
		//包装错误响应给前端
		l.Error(err)
		err = c.WriteMessage(typ, wrapMsg(FAIL, "服务掉用失败"))
		if err != nil {
			l.Error(err)
		}
		return
	}
	// coze成功响应
	for {
		event, err := resp.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) {
				err = c.WriteMessage(typ, wrapMsg(FAIL, "对话超时"))
			} else if errors.Is(err, io.EOF) {
				err = c.WriteMessage(typ, wrapMsg(END, "回答结束"))
			} else {
				l.Error(err)
				err = c.WriteMessage(typ, wrapMsg(FAIL, "服务中断"))
			}
			if err != nil {
				l.Error(err)
			}
			return
		}
		if event.Event == coze.ChatEventConversationMessageDelta {
			err = c.WriteMessage(typ, wrapMsg(SUCCESS, event.Message.Content))
			if err != nil {
				l.Error(err)
			}
		}
	}
}

func wrapMsg(code int, msg string) []byte {
	return []byte(fmt.Sprintf("{\"code\": \"%d\", \"msg\": \"%s\"}", code, msg))
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
