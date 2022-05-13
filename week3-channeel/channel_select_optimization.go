package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//需要初始化随机数的资源库，如果不执行这一行，不管运行多少次都是返回相同的值
	rand.Seed(time.Now().UnixNano())
	no := rand.Intn(6)
	fmt.Printf("[no is %d, type is %T]\n", no, no)
	no *= 1000
	du := time.Duration(int32(no)) * time.Millisecond
	fmt.Println("timeout duration is :", du)
	wg.Done()
	if isTimeout(&wg, du) {
		fmt.Println("Time out")
	} else {
		fmt.Println("Not time out  ")
	}

}


func isTimeout(wg *sync.WaitGroup, du time.Duration) bool {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		defer close(ch1)
		wg.Wait()
	}()

	select {
	case <- ch1:
		return false
	case <- time.After(du):
		return true

	}
}
//这是一种模拟，让程序的等待时间可以更具传入参数进行不同的超时间长判断， 这里是用随机数来模拟时长的，真实的项目中可以根据配置参数或者统计参数在运行时传递到函数中。
