package public

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

type Content struct {
	Tid int
	Context string
    Status_code int
    Url string
}


func (self *Content) Get (tid int, result *Content) error {
    result.Tid = tid
    url := fmt.Sprintf("http://club.autohome.com.cn/bbs/thread-c-442-%d-1.html", tid)
    
    // 新建请求客户端
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, strings.NewReader(""))

    req.Header.Set("User-Agent", "")
    req.Header.Set("X-Forwarded-For", "123.125.71.97")
    
    // 发送请求
    r, _ := client.Do(req)

    defer r.Body.Close()
    text, _ := ioutil.ReadAll(r.Body)
    result.Context = string(text)
    result.Status_code = r.StatusCode
    result.Url = url
    return nil
}