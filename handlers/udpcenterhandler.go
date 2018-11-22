package handlers

import (
	"hzj/comm/udp"
	"encoding/json"
)

func ReadUdpMessage(){
	for {
		data := udp.ReadMessage()

		dd, err := json.Marshal(data.Data)
		if err == nil {
			if hd, ex := UdpActionsMapping[data.Action]; ex {
				hd.Run(string(dd))
			}

		}
	}
}
