package demo

import (
	"fmt"
	"github.com/jponge/vertx-go-tcp-eventbus-bridge/eventbus"
	"gox/etcd"
	"gox/kv"
	"gox/random"
	"log"
	"testing"
	"time"
)

// TestAdd 单元测试
//func TestAdd(t *testing.T) {
//	if ans := add(1, 2); ans != 3 {
//		t.Error("add(1, 2) should be equal to 3")
//	}
//}

// BenchmarkFib 压力测试（必须以BenchmarkFib开头） go test -bench="Fib$" -benchtime=5s -count=3
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}

}

// 测试slice
func TestFunc(t *testing.T) {
	var ans map[int][]int
	ans = make(map[int][]int)
	ans[0] = []int{1}

	examNo := 1

	for studentno, exams := range ans {
		for _, okExamNo := range exams {
			if okExamNo == examNo {
				// 删除指定位置的切片
				exams = append(exams[:examNo-1], exams[examNo:]...)
				ans[studentno] = exams
			}
		}
	}

	fmt.Printf("xx:%v", ans)

}

// TestKv 测试缓存
func TestKv(t *testing.T) {
	kv.Set("str", "abc")
	fmt.Println(kv.Get("str"))
}

// TestDB 测试数据库连接。使用sqlx
//func TestDB(t *testing.T) {
//}

func TestEtcd(t *testing.T) {
	etcd.Set("sunqb1", "中国人1")
	etcd.SetExpire("sunqb2", "中国人2", 10)
	fmt.Println(etcd.Get("sunqb1"))
	fmt.Println(etcd.Get("sunqb2"))
}

func TestGet(t *testing.T) {
	fmt.Println(etcd.Get("sunqb2"))
}

func TestRandom(t *testing.T) {
	fmt.Println(random.GetRandomString(16))
}

// TestVertx 测试 vertx-go-tcp-eventbus-bridge
func TestVertx(t *testing.T) {
	stop := make(chan int, 1)

	// 获取一个eventbus连接
	eventBus, err := eventbus.NewEventBus("localhost:7000")
	if err != nil {
		// 打印错误日志并终止执行程序
		log.Fatal("Connection to the Vert.x bridge failed: ", err)
	}

	// 获取一个dispatcher来持续接收messages
	dispatcher := eventbus.NewDispatcher(eventBus)
	dispatcher.Start()

	// 在地址sample.clock.ticks 获取一个长度为8的队列(chan)
	ticks, id, err := dispatcher.Register("sample.clock.ticks", 8)
	if err != nil {
		log.Fatal("Registration on sample.clock.ticks failed: ", err)
	}
	log.Println("sample.clock.ticks channel registered with ID=", id)

	// 消费消息-方式1：开一个新协程，在里面运行匿名的函数。
	go func() {
		for {
			// 从ticks里面读取数据，不用变量存储，就是玩。
			reply := <-ticks
			log.Println("[tick]:", reply)
		}
	}()

	// 消费消息-方式2
	//go func() {
	//	for reply := range ticks {
	//		log.Println("Echo, got back: ", reply)
	//	}
	//}()

	// 生产消息-方式1
	go func() {
		for {
			// 这一行每3s生产一次消息
			<-time.After(time.Second * 3)
			eventBus.SendWithReplyAddress("sample.echo", "sample.echo.reply", nil, map[string]string{
				"source": "Go Client",
				"what":   "A tick!",
			})
		}
	}()

	// 生产消息-方式2
	//go func() {
	//	for {
	//		select {
	//		case <-time.After(time.Second * 3):
	//			eventBus.SendWithReplyAddress("sample.echo", "sample.echo.reply", nil, map[string]string{
	//				"source": "Go Client",
	//				"what":   "A tick!",
	//			})
	//		}
	//	}
	//}()

	// 从停止管道里面读取消息判断程序是否做停止操作，需要程序员自己做
	<-stop
}
