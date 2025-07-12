package main

import (
	"Debate-System/utils/jwtx"
	"Debate-System/utils/syncx"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Hub struct {
	// syncx.Map 是我对 sync.Map 的一个简单封装
	conns *syncx.Map[int64, *websocket.Conn]
}

func (h *Hub) AddConn(user_id int64, c *websocket.Conn) {
	h.conns.Store(user_id, c)
	go func() {
		for {
			typ, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNoStatusReceived) {
					//客户端断开连接
					h.conns.Delete(user_id)
				} else {
					log.Println(err)
				}
				return
			}
			// 开始转发
			h.conns.Range(func(key int64, value *websocket.Conn) bool {
				if key == user_id {
					return true
				}
				err = value.WriteMessage(typ, message)
				if err != nil {
					log.Println(err)
				}
				// 返回 true，确保会继续往后遍历
				return true
			})
		}
	}()
}

func main() {
	upgrader := websocket.Upgrader{}
	hub := &Hub{conns: &syncx.Map[int64, *websocket.Conn]{}}
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		// 没有额外的 header 的设置
		user_id, err := jwtx.ParseToken("app-api-secret", request)
		if err != nil {
			log.Println(err)
			writer.Write([]byte("用户信息异常，请重新登陆"))
			return
		}
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println(err)
			writer.Write([]byte("初始化 websocket 失败"))
			return
		}
		hub.AddConn(user_id, conn)
		fmt.Println("用户加入成功", user_id)
	})
	http.ListenAndServe(":8080", nil)
}
