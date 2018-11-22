package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HttpCenter(c *gin.Context) {
	value, exist := c.GetQuery("data")
	if !exist {
		log.Println("data >> is nil")
	}

	pd := struct {
		Action int `json:"action"`
		Data interface{} `json:"data"`
	}{}

	err := json.Unmarshal([]byte(value), &pd)

	var bak interface{}

	code := 0
	errmsg := "成功"
	if err != nil {
		code = 1
		errmsg = "未知消息"
	} else {

		pdata := "{}"
		if pd.Data != nil {
			pj, err := json.Marshal(pd.Data)
			if err == nil {
				pdata = string(pj)
			}
		}

		if hd, ex := HttpActionsMapping[pd.Action]; ex {
			code, bak = hd.Run(pdata)
		}

	}

	var bak_data string
	if bak != nil {
		bakstr, _ := json.Marshal(bak)
		bak_data = string(bakstr)
	}

	c.JSON(http.StatusOK, gin.H{
		"action":pd.Action,
		"code": code,
		"errmsg": errmsg,
		"data":bak_data })
	return
}
