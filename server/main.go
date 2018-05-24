package main

import (
	"strings"
    "fmt"
    "net"
    "net/http"
    "net/rpc"
    "io/ioutil"
    "flag"
)

func input() string {
	port := flag.String("port", "0.0.0.0:1234", "port")
	flag.Parse()
	return *port
}


type Content struct {
	Tid int
	Context string
	Status_code int
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
    return nil
}



func main() {

    mirror := new(Content)
    rpc.Register(mirror)
    rpc.HandleHTTP()
    port := input()

    l, err := net.Listen("tcp", port)
    if err != nil {
        fmt.Println("监听失败，端口可能已经被占用")
        }
        fmt.Println("正在监听1234端口")
        http.Serve(l, nil)
}


