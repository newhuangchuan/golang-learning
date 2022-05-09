package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(2 * time.Second)

}

func writeChan(c chan int, x int) {
	log.Println(x)
	c <- x
	close(c)
	log.Println(x)
}
