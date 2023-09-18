package main

import "fmt"

func panic_in_defer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("wuwuuwuw",err)
		}else {
			fmt.Println("hahhahaha")
		}

	}()
	defer func() {
		panic("defer inner panic")
	}()

	panic("function inner panic")
}

func main() {
	panic_in_defer()
}

/*
- panic 仅有最后一个可以被recover捕获
- 这个异常将会覆盖掉main中的异常，最后这个异常被第二个执行的defer捕获到了
*/
