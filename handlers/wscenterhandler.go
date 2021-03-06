package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"hzj/comm/wsi"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

func WsCenter(c *gin.Context) {
	var(
		wsConn *websocket.Conn
		err error
		conn *wsi.Connection
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn , err = upgrader.Upgrade(c.Writer,c.Request,nil); err != nil{
		return
	}

	if conn , err = wsi.InitConnection(wsConn); err != nil{
		goto ERR
	}

	// 启动心跳
	go func() {

		if hd, ex := SystemActionMapping[ConnHeartAction]; ex {
			hd.Run("", conn)
		}

	}()

	for {

		if data , err := conn.ReadMessage();err != nil{
			goto ERR
		} else {

			jdata := "{}"
			if data.Data != nil {
				dj, err := json.Marshal(data.Data)
				if err == nil {
					jdata = string(dj)
				}
			}
			if hd, ex := ActionsMapping[data.Action]; ex {
				hd.Run(jdata, conn)
			} else {
				log.Println("不存在的action:", data.Action)
				goto ERR
			}
		}
	}
ERR:
	conn.Close()
	if hd, ex := SystemActionMapping[ConnCloseAction]; ex {
		hd.Run(conn.Uid, nil)
	}
}
