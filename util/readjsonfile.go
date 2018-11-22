package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func Parsejf(jfile string, v interface{}){

	data, err := ioutil.ReadFile(jfile)
	if err != nil {
		log.Println(jfile, ">>读取失败>>", err)
		return
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		log.Println(jfile, ">>解析失败>>", err)
	}

}