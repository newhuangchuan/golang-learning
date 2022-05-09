package main

import (
	"log"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		log.Println("can read from ch1")
	case <-ch2:
		log.Println("can read from ch2")
	default:
		log.Println("defaults......")

	}

}
