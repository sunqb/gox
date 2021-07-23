package demo

import (
	"errors"
	"fmt"
	_ "net/http/pprof"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// channel
var ch = make(chan string, 10)

/**
加法
*/
func add(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return
}

/**
测试自定义错误使用errors
*/
func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error:name is null")
	}
	fmt.Println("hello,", name)
	return nil
}

/**
异常
*/
func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

/**
下载
*/
func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}

func init() {

}

func init() {

}

// 常量
const (
	n1 = 1
	n2 = 2
	_
	n3 = 4
	n4 = iota
	n5
)

// 定义类型
type MyInt int

// 给MyInt类型添加SayHello方法
func (m MyInt) SayHello() {

}

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func demo() {
	// 1、示例
	//acd := "2"
	//fmt.Println("1+1=", acd)

	// 2、变量
	//var a0 int
	//var a1 int = 1
	//var a2 = 1
	//a3 := 1
	//msg0 := "Hello World!"

	// 3、简单类型
	//var a int8 = 10
	//var c1 byte = 'a'
	//var b float32 = 12.2
	//var msg = "H"
	//ok := false

	// 4、字符串
	//str1 := "Golang"
	//str2 := "Go语言"
	//fmt.Println(reflect.TypeOf(str2[2]).Kind())
	//fmt.Println(str1[2], string(str1[2]))
	//fmt.Printf("%d %c\n", str2[2], str2[2])

	// 5、数组
	//c2 := float32(5.3)
	//fmt.Println(c2)
	//arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr)
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	// 6、切片
	//slice1 := make([]float32, 3)
	//slice1 = append(slice1, 1, 2, 3, 4)
	//fmt.Println(slice1)
	//fmt.Println(reflect.TypeOf(slice1))

	// 7、字典：标准使用make
	//m1 := make(map[string]int)
	//m2 := map[string]string{
	//	"sam":   "male",
	//	"alice": "female",
	//}
	//m1["sam"] = 18
	//m2["susan"] = "f"
	//fmt.Println(m1)
	//fmt.Println(m2)

	// 8、指针
	//strp := "Golang"
	//var p1 *string = &strp
	//*p1 = "Hello"
	//fmt.Println(strp)
	//
	//str := new(string)
	//*str = "welcome"
	//fmt.Println(*str)

	// 9、if
	//if age := 18; age < 18 {
	//	fmt.Println("kid")
	//} else {
	//	fmt.Println("adult")
	//}

	// 10、switch
	//type Gender int8
	//const
	//(
	//	MALE   Gender = 1
	//	FEMALE Gender = 2
	//)
	//
	//gender := MALE
	//switch gender {
	//case FEMALE:
	//	fmt.Println("female")
	//case MALE:
	//	fmt.Println("male")
	//default:
	//	fmt.Println("unknown")
	//}

	// 11、for
	// fori
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	if sum > 50 {
	//		break
	//	}
	//	sum += i
	//}

	// forr
	//nums := []int{10, 20, 30, 40}
	//for i, num := range nums {
	//	fmt.Println(i, num)
	//}
	//m22 := map[string]string{
	//	"Sam":   "Male",
	//	"Alice": "Female",
	//}
	//for key, value := range m22 {
	//	fmt.Println(key, value)
	//}

	// 12、函数。可以有多返回值、且返回变量可提前定义
	//fmt.Println(add(1, 2))

	// 13、异常
	//_, err := os.Open("filename.txt")
	//if err != nil {
	//	fmt.Printf("error is:[ %s ]", err)
	//}
	//fmt.Println()

	// 自定义错误
	//err1 := hello("")
	//if err1 != nil {
	//	fmt.Println(err1)
	//}

	// 异常捕获
	//fmt.Println("")

	// channel.注意这里有类型转换的坑。即int怎么转换为string。建议使用Itoa函数
	//for i := 0; i < 3; i++ {
	//	go download("a.com/" + strconv.Itoa(i))
	//}

	// 单元测试
	// 见demo_test.go

	// Gin web 框架
	//r := gin.Default()
	//r.GET("/user/:name", func(c *gin.Context) {
	//	c.String(http.StatusOK,"Htllo, Geektutu")
	//})
	//r.GET("/users", func(c *gin.Context) {
	//	name := c.Query("name")
	//	role := c.DefaultQuery("role", "teacher")
	//	c.String(http.StatusOK, "%s is a %s", name, role)
	//})
	//r.Run()

	// init函数和main函数
	//var a int
	//a = 1
	//a += 2
	//fmt.Println(a)
	//fmt.Println(n4)

	// strings内置函数使用
	//a := "中先"
	//strings.Contains(a,"中")

	// 指针的使用：指针类型、指针地址、指针取值
	//num := 1
	//fmt.Println("num的地址是：", &num)
	//
	//ptr := &num
	//*ptr += 1
	//
	//fmt.Println("修改后num的值是：", *ptr)

	//fmt.Println(time.After(time.Second*3))

	// Goexit() 直接退出协程
	//go func() {
	//	defer fmt.Println("A.defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		// 结束协程
	//		runtime.Goexit()
	//		defer fmt.Println("C.defer")
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//}()
	//for {
	//}

	// Gosched() 切一下协程
	//go func(s string) {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println(s, i)
	//	}
	//}("world")
	//// 主协程
	//for i := 0; i < 2; i++ {
	//	// 切一下，再次分配任务
	//	runtime.Gosched()
	//	fmt.Println("hello", i)
	//}

	// channel。下面这段程序直接在这里执行会报死锁的错误
	//var channel chan string
	//channel = make(chan string, 1)
	//
	//channel <- "afu" // 向通道里发送字符串
	//
	//x := <-channel // 从channel中接收值并赋值给x
	//
	//fmt.Println(x)

	// channel 练习
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//// 开启goroutine将0~100的数发送到ch1中
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	//go func() {
	//	for {
	//		i, ok := <-ch1 // 通道关闭后再取值ok=false
	//		if !ok {
	//			break
	//		}
	//		ch2 <- i * i
	//	}
	//	close(ch2)
	//}()
	//// 在主goroutine中从ch2中接收值打印
	//for i := range ch2 { // 通道关闭后会退出for range循环
	//	fmt.Println(i)
	//}

	// 需要2个管道
	// 1.job管道
	//jobChan := make(chan *Job, 128)
	//// 2.结果管道
	//resultChan := make(chan *Result, 128)
	//// 3.创建工作池
	//createPool(64, jobChan, resultChan)
	//// 4.开个打印的协程
	//go func(resultChan chan *Result) {
	//	// 遍历结果管道打印
	//	for result := range resultChan {
	//		fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
	//			result.job.RandNum, result.sum)
	//	}
	//}(resultChan)
	//var id int
	//// 循环创建job，输入到管道
	//for {
	//	id++
	//	// 生成随机数
	//	r_num := rand.Int()
	//	job := &Job{
	//		Id:      id,
	//		RandNum: r_num,
	//	}
	//	jobChan <- job
	//}

	// 定时器Timer：延迟执行
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//// create a nobuf channel and a goroutine `timer` will write it after 2 seconds
	//timeAfterTrigger := time.After(time.Second * 10)
	//// will be suspend but we have `timer` so will be not deadlocked
	//curTime, _ := <-timeAfterTrigger
	//// print current time
	//fmt.Println(curTime.Format("2006-01-02 15:04:05"))
	//
	//
	//// 创建一个计时器：定时执行
	//timeTicker := time.NewTicker(time.Second * 2)
	//i := 0
	//for {
	//	if i > 5 {
	//		break
	//	}
	//
	//	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//	i++
	//	<-timeTicker.C
	//
	//}
	//// 清理计时器
	//timeTicker.Stop()

	// select多路复用
	// 没有使用select会怎么做
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//data, ok := <-ch1
	//data, ok := <-ch2
	//
	//select {
	//case <-ch1:
	//// ch1成功读到数据
	//case ch2 <- 1:
	//default:
	//
	//}

	// 并发sync包
	// sync.Once 只执行一次。就算手动调用多次，也只会调用一次。
	//var once sync.Once
	//funO := func() {
	//	fmt.Println("a", "b", "c")
	//}
	//once.Do(funO)
	//once.Do(funO)

	// sync.Map，安全的map

	// atomic。atomic的cas底层实现逻辑？
	//atomic.CompareAndSwapInt32()

	//res := concatString("a", "b", "c")
	//fmt.Println(res)

	// slice和数组互相转换
	// 数组转换为切片
	//arr := [2]string{"a","b"}
	//slice := arr[:]
	//arr[0] = "c"
	//fmt.Println(slice)
	//fmt.Println(arr)
	//fmt.Println(reflect.TypeOf(slice))

	// 切片转换为数组
	//slice := []byte("abcdefgh")
	//var arr [4]byte
	//copy(arr[:], slice[:4])
	//fmt.Println(arr)
	//fmt.Println(reflect.TypeOf(arr))

	//var studentPool = sync.Pool{
	//	New: func() interface{
	//		return new(Student)
	//	},
	//}

	// 时间处理
	//fmt.Println(time.Now().Unix())	// 1624600901
	//fmt.Println(time.Time{}.Unix()) // -62135596800
	//fmt.Println(time.Now().Format(time.RFC3339)) // 2021-06-25T14:01:41+08:00

	// 定时调用
	//worker()

	// 阻塞多少时间
	//select {
	//case <-time.After(1000000 * time.Second):
	//	fmt.Println(1)
	//}

	// channel控制并发数。：先要从 limit 中拿“许可证”，拿到许可证之后，才能执行 w()，并且在执行完任务，要将“许可证”归还。这样就可以控制同时运行的 goroutine 数。
	//var limit = make(chan int, 3)
	//for _, w:range work{
	//	go func (){
	//	limit <- 1
	//	w()
	//	<- limit
	//}()
	//}

	// channel, happen before semantics
	//go aGoroutine()
	//done <- true
	//println(msg)

	//var slice1 []string
	//
	//fmt.Printf("%p",&slice1)
	//fmt.Println()
	//
	//slice1 = append(slice1,"1")
	//fmt.Printf("%p",&slice1)
	//
	//fmt.Println(slice1)

	// 测试 WaitGroup
	//testWaitGroup()

	//cha1 := make(chan string)
	//cha2 := make(chan string)
	//for {
	//	select {
	//	case e := <-cha1:
	//		fmt.Println("cha1:", e)
	//	case e := <-cha2:
	//		fmt.Println("cha2", e)
	//	default:
	//		fmt.Println("default")
	//	}
	//}
	//for i2 := range cha1 {
	//
	//}
}

// 使用builder拼接字符串
func concatString(str ...string) string {
	var builder strings.Builder
	for _, s := range str {
		builder.WriteString(s)
	}
	arr := [3]string{"1", "2", "3"}
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(arr))
	return builder.String()
}

func worker() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}

var msg string
var done = make(chan bool)

func aGoroutine() {
	msg = "H W"
	<-done
}

func testWaitGroup() {
	var wg sync.WaitGroup
	var counter int32
	n := 10
	for i := 0; i < n; i++ {
		wg.Add(1) // 这个要放在外面
		go func(i int) {
			atomic.AddInt32(&counter, 1)
			// counter += 1 // Dangerous!!
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
