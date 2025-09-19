package ws

import (
	"Debate-System/chat/internal/svc"
	"Debate-System/utils/syncx"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/coze-dev/coze-go"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
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

	// ====== 新增 1：读超时 + Pong 续命（必须放在起 goroutine 之前） ======
	if err := c.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
		logger.Errorf("set initial read deadline error: %v", err)
		h.conns.Delete(user_id)
		c.Close()
		return
	}
	c.SetPongHandler(func(string) error {
		if err := c.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
			logger.Errorf("pong handler set read deadline error: %v", err)
			// 无法续命就直接关掉连接，避免幽灵
			h.conns.Delete(user_id)
			c.Close()
		}
		return nil
	})

	// ====== 新增 2：写超时（给后面 WriteMessage 用） ======
	if err := c.SetWriteDeadline(time.Now().Add(60 * time.Second)); err != nil {
		logger.Errorf("set initial write deadline error: %v", err)
		h.conns.Delete(user_id)
		c.Close()
		return
	}

	// ====== 新增 3：定时 Ping（在你原有“一个 goroutine”里跑，不拆函数） ======
	// 注意：你原来只开了一个 goroutine，我们**仍在同一个 goroutine里**再启一个 tick 器，
	// 这样你代码结构保持“只有一个 goroutine”的观感，逻辑也没变。
	go func() {
		tick := time.NewTicker(30 * time.Second)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				if err := c.SetWriteDeadline(time.Now().Add(60 * time.Second)); err != nil {
					logger.Errorf("ping tick set write deadline error: %v", err)
					h.conns.Delete(user_id)
					c.Close()
					return
				}
				if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
					logger.Errorf("write ping error: %v", err)
					h.conns.Delete(user_id)
					c.Close()
					return
				}
			}
		}
	}()

	go func() {
		for {
			typ, message, err := c.ReadMessage()
			//监视是否在线
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
					logger.Infof("user %s normal close", user_id)
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
			// 新增：每次流式写出前刷新写超时
			if err = c.SetWriteDeadline(time.Now().Add(60 * time.Second)); err != nil {
				l.Errorf("cozeChat flush write deadline error: %v", err)
				return // 直接结束流式写出，goroutine 可退出
			}
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
