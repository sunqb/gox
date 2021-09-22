package demo

import (
	"fmt"
	"testing"
)

type Emplyoee struct {
	Id   int
	Name string
}

type Pro interface {
	hello() string
}

func (god *Emplyoee) hello() string {
	fmt.Println("12")
	return ""
}

func TestFirstTry(t *testing.T) {
	//t.Log("f t")
	//str := "中"
	//
	//str1 := []rune(str)
	//t.Log(len(str))
	//t.Log(len(str1))

	//e1 := Emplyoee{1, "2"}
	//e2 := &Emplyoee{1, "2"}
	//e3 := new(Emplyoee)
	//
	//t.Logf("%T", e1)
	//t.Logf("%T", e2)
	//t.Logf("%T", e3)

	//var p *Emplyoee
	//p = new (Emplyoee)
	//p.hello()
	//
	//fmt.Println(p)
	//
	//var pd interface{}
	//if i,ok := pd.(int);ok{
	//	fmt.Println(i)
	//}

	//var mut sync.Mutex
	//counter := 0 // 多个协程多共享变量的操作需要加锁
	//for i := 0; i < 5000; i++ {
	//	go func() {
	//		defer func() {
	//			mut.Unlock()
	//		}()
	//
	//		mut.Lock() // 这个锁保证了对counter的++都是期望的值。如果没有锁，counter的值会错乱的更多。因为有很多协程的计算会不按照期望值计算。
	//		counter++
	//	}()
	//}
	//time.Sleep(1 * time.Second) // 这个去掉后counter就不是5000了
	//t.Logf("counter=%d", counter)

	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 会阻塞，直到retCh有结果返回

}

// 模拟java的Future
func AsyncService() chan string {
	retCh := make(chan string)

	go func() {
		ret := service()
		retCh <- ret
	}()
	return retCh
}

func service() string {
	return "service"
}

func otherTask() {
	fmt.Println("other task")
}

// FirstResponse 多个任务，只要任意一个返回了，函数就返回。
func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string) // 这里使用了非缓存的channel，会导致协程泄露，因为chan在第一个返回后，其他协程的会一致阻塞，这里只要改为make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

// runTask 一个简单接收参数的任务
func runTask(i int) string {
	fmt.Println(i)
	return fmt.Sprintf("i is %v", i)
}
