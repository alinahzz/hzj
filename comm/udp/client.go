package udp

import (
	"encoding/json"
	"log"
	"net"
)

func Push(addr string, data []byte) {

	if addr == "" {
		return
	}

	saddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		log.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, saddr)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("udpout>>", string(data))
}

const SYNCMESSAGE = 666
func PushSyncMessage(saddr string, uid string, data interface{}) {

	if uid == "" {
		return
	}

	if saddr == "" {
		return
	}

	pp := make(map[string]interface{})
	mm := make(map[string]interface{})
	pp["action"] = SYNCMESSAGE
	pp["data"] = mm
	mm["uid"] = uid
	mm["data"] = data

	bb, _ := json.Marshal(pp)
	Push(saddr, bb)

}

const SYNCREQUST = 888
func PushSyncRequest(saddr string, suba int, data interface{}) {

	if saddr == "" {
		return
	}

	pp := make(map[string]interface{})
	mm := make(map[string]interface{})
	pp["action"] = SYNCREQUST
	pp["data"] = mm
	mm["action"] = suba
	mm["data"] = data

	bb, _ := json.Marshal(pp)
	Push(saddr, bb)

}
