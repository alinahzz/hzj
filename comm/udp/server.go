package udp

import (
	"encoding/json"
	"log"
	"net"
)

var InChan chan *Req

type Req struct {
	Action int `json:"action"`
	Data interface{} `json:"data"`
}

func init() {
	InChan = make(chan *Req, 1000)
}

func ReadMessage() (data *Req) {
	data = <- InChan
	return
}

func Start(saddr string){

	log.Println("udp-path:>>", saddr)

	addr, err := net.ResolveUDPAddr("udp", saddr)
	if err != nil {
		log.Println(">>udp server start faild!!!")
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(">>udp conn get faild!!!", err)
		return
	}

	defer conn.Close()

	data := make([]byte, 2048)
	for {
		dl, _, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
			continue
		}

		var req Req
		err = json.Unmarshal(data[:dl], &req)
		if err != nil {
			log.Println(err)
			continue
		}

		InChan <- &req

		log.Println("udpin<<", req)
	}
}