package main

import (
	"github.com/golang-learning/interface/tty"
	"log"
)

var a = tty.Test{
	X: 18,
}

func main() {

	log.Println(a.X)

}
