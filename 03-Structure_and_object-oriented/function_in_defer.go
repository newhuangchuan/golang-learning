package main

import "fmt"

func function_in_defer(index int, values int) int {
	fmt.Println(index)
	return index
}

func main() {
	defer function_in_defer(1,function_in_defer(3,0))
	defer function_in_defer(2,function_in_defer(4,0))
}


/*
- 先defer函数压栈，第一个defer
- 先计算参数的值3，然后压栈，
- 执行函数得到4，作为参数，进行压栈
*/