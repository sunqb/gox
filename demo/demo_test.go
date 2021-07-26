package demo

import (
	"fmt"
	"gox/etcd"
	"gox/kv"
	"gox/random"
	"testing"
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
