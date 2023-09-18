package main

import "fmt"

func main()  {
	defer_func()

}


func defer_func() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("被recover 捕获到的panic信息 内容是： ", err)
		}
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

	panic("error......")
}
