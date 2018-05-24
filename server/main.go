package main

import (
    "fmt"
    "net"
    "net/http"
    "net/rpc"
    "flag"
    "spider/public"
)

func input() string {
	port := flag.String("port", "0.0.0.0:1234", "port")
	flag.Parse()
	return *port
}


func main() {

    mirror := new(public.Content)
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


