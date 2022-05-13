package main

import (
	"fmt"
	"math/rand"
)

//channel 可以连接goroutine，如果一个goroutine的输出是拎一个goroutine的输入，就叫做pipeline。可以理解为pipeline是虚拟的，用来连接goroutine和channel，并且最终形成一个goroutine的输出成为另一个goroutine的输入，且是使用channel传递数据的。
//使用pipeline的好处有三点：
//首先，在程序中形成一个清晰稳定的数据流，我们在使用的时候不需要过多考虑gorouine和channel的相关通信和状态问题。
//其次，在一个pipeline内，不需要把数据再保存为变量，节省了内存空间并提高了效率。
//最后，使用pipeline能够简化程序的复杂度，便于维护
var done = false
var Mess = make(map[int]bool)

func main() {
	A := make(chan int)
	B := make(chan int)
	go sendRan(50, 10, A)
	go receive(B, A)
	sum(B)
}

func genRandom(max, min int) int {
	return rand.Intn(max - min) + min
}

func sendRan(max, min int, out chan int)  {
	for {
		if done {
			close(out)
			return
		}
		out <- genRandom(max, min)
	}
}

func receive(out chan<- int, in <-chan int) {
	for r := range in {
		fmt.Println(" " , r)
		_, ok := Mess[r]
		if ok {
			fmt.Println("duplicate num is : ", r)
			done = true
		} else {
			Mess[r] = true
			out <- r
		}
	}
	close(out)
}

func sum(in <-chan int) {
	var sum int
	for i := range in {
		sum += i
	}
	fmt.Println("The sum is : ", sum)
}
