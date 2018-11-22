package wsi

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Connection struct{
	wsConnect *websocket.Conn
	inChan chan *Req
	outChan chan *Rsp
	closeChan chan byte

	mutex sync.Mutex  // 对closeChan关闭上锁
	isClosed bool  // 防止closeChan被关闭多次

	Uid string
}

type Req struct {
	Action int `json:"action"`
	Data interface{} `json:"data"`
}

type Rsp struct {
	Action int `json:"action"`
	Code int `json:"code"`
	Data map[string]interface{} `json:"data"`
}

func InitConnection(wsConn *websocket.Conn)(conn *Connection ,err error){
	conn = &Connection{
		wsConnect:wsConn,
		inChan: make(chan *Req,10),
		outChan: make(chan *Rsp,50),
		closeChan: make(chan byte,1),

	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}

func (conn *Connection)ReadMessage()(data *Req , err error){

	select{
	case data = <- conn.inChan:
	case <- conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

func (conn *Connection)WriteMessage(data *Rsp)(err error){

	select{
	case conn.outChan <- data:
	case <- conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

func (conn *Connection)Close(){
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection)readLoop(){
	var(
		data Req
		err error
	)
	for{
		if err = conn.wsConnect.ReadJSON(&data); err != nil{
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲位置
		select{
		case conn.inChan <- &data:
		case <- conn.closeChan:		// closeChan 感知 conn断开
			goto ERR
		}

		log.Println("In>>",data)

	}

ERR:
	conn.Close()
}

func (conn *Connection)writeLoop(){
	var(
		data *Rsp
		err error
	)

	for{
		select{
		case data= <- conn.outChan:
		case <- conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteJSON(*data); err != nil{
			goto ERR
		}
		log.Println("Out<<",*data)
	}

ERR:
	conn.Close()

}
