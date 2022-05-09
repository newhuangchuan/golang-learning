package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan  string)
	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "ch1 is ready!"
	}()

	select {
	case mess := <-ch1:
		fmt.Println(mess)
	case t := <- time.After(2 * time.Second):
		fmt.Println("ch1 timeout ", t)
		
	}

	ch2 := make(chan string)
	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "ch2 is ready !"
	}()
	select {
	case mess := <- ch2:
		fmt.Println(mess)
	case t := <-time.After(5 * time.Second):
		fmt.Println("ch2 is timeout", t)
	}

}
