package main

import "fmt"

func main() {
	panicIndefer()


}

func panicIndefer()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("hahahha")
		}
	}()

	defer func() {
		panic("defer 内部触发的一个panic")
	}()

	//外部触发一个panic
	panic("外部触发的panic")
}
