package comm

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)


func StartWebSocket(tsl bool, wspath string, httppath, addr string, method string,
	httphandler gin.HandlerFunc, wshandler gin.HandlerFunc, certfile string, keyfile string) {

	log.Println("http-path:>>", addr, httppath)
	log.Println("websocket-path:>>", addr, wspath)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	if httppath != "" {
		switch method {
		case "GET":
			router.GET(httppath, httphandler)
			break
		case "POST":
			router.POST(httppath, httphandler)
			break
		case "GET&POST":
			router.GET(httppath, httphandler)
			router.POST(httppath, httphandler)
			break
		}
	}

	if wspath != "" {
		router.GET(wspath, wshandler)
	}

	if tsl {
		http.ListenAndServeTLS(addr, certfile, keyfile, router)
	} else {
		http.ListenAndServe(addr, router)
	}

}


func HttpGet(url string) string{

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	bak := string(body)
	log.Println(bak)

	return bak
}

func HttpPost(addr string, params map[string] string) string{

	val := url.Values{}

	for k,v := range params {
		val.Set(k, v)
	}

	pv := val.Encode()
	resp, err := http.Post(addr, "application/x-www-form-urlencoded",
		strings.NewReader(pv))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	bak := string(body)
	log.Println(bak)

	return bak
}
