package handlers

import "hzj/comm/wsi"

type Actionhandler interface {
	Run(data string, conn *wsi.Connection)
}

type Actionhandlerimp func(data string, conn *wsi.Connection)

func (f Actionhandlerimp) Run(data string, conn *wsi.Connection) {
	f(data, conn)
}

type Udpactionhandler interface {
	Run(data string)
}

type Udpactionhandlerimp func(data string)

func (f Udpactionhandlerimp) Run(data string){
	f(data)
}

type Httpactionhandler interface {
	Run(data string) (int, interface{})
}

type Httpactionhandlerimp func(data string) (int, interface{})

func (f Httpactionhandlerimp) Run(data string) (int, interface{}) {
	return f(data)
}

/** 消息协议与处理映射 */
var ActionsMapping map[int]Actionhandler

/** 服务器间协议与处理映射 */
var UdpActionsMapping map[int]Udpactionhandler

/** HTTP协议与处理映射 */
var HttpActionsMapping map[int]Httpactionhandler

const ConnCloseAction = 444

func init() {
	ActionsMapping = make(map[int]Actionhandler)
	UdpActionsMapping = make(map[int]Udpactionhandler)
	HttpActionsMapping = make(map[int]Httpactionhandler)
}
