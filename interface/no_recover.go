package main

import "fmt"

func main() {
	defer_func()

}

func defer_func() {

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



	/*
	3
	2
	1
	panic: I am panic
	 */
}
