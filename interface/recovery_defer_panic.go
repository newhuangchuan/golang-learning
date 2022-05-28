package main

import "fmt"

func main() {
	defer_func_panic()

}

func defer_func_panic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[被recover捕获的panic  内容是 ： %v]\n", err)
		}
		fmt.Println("0")
	}()
	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()
	panic("I am panic")
}
