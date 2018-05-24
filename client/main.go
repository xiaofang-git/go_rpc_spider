package main
import (
	"fmt"
	"net/rpc"
	"sync"
	"time"
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	"spider/public"

)

var wg sync.WaitGroup

var work_list = [...]string{
	"127.0.0.1:1234",
	// "218.60.41.200:1234",
	// "218.60.41.200:1234",
}

func id_db() int {
	// 从数据库获取最大id
	return 30389
}

func id_api() int {
	通过api获取最大id
	conn, _ := gorose.Open(public.DbConfig)
	key,_ := conn.Table("mirror").Max("mirror_key")
	fmt.Println(key)
	return 30389
}


var ch = make(chan public.Content, 10)

func task(ip string, tid int)  {
	// 获取结果
	client, err := rpc.DialHTTP("tcp", ip)
	if err != nil {
		fmt.Println("链接rpc服务器失败:", err)
		}
		result := new(public.Content)

		err = client.Call("Content.Get", tid, &result)

		if err != nil {
			fmt.Println("调用远程服务失败", err)

			}

		select {
			case ch <- *result:
			case <- time.After(time.Second * 5):
		}
		wg.Done()
}


func worker()  {

	// // 等待所有任务执行完毕
	start := id_db()
	end := id_api()
	for tid:= start; tid<=end; tid++ {
		time.Sleep(time.Second*1)
		// 遍历需要抓取的id列表
		w := work_list[tid % len(work_list)]
		wg.Add(1)
		go task(w, tid)
	}

}

func insert() {
	start := id_db()
	end := id_api()
	conn, _ := gorose.Open(public.DbConfig)
	for i:=start; i<=end; i++ {
		mirror := <- ch
		_, err := conn.Table("mirror").Data(map[string]interface{}{
			"mirror_key": mirror.Tid,
			"mirror_status": mirror.Status_code,
			"mirror_url": mirror.Url,
			}).Insert()
		fmt.Println(err)
	}

	conn.Close()
	wg.Done()
}


func main() {

	worker()
	wg.Add(1)
	go insert()
	wg.Wait()
}
