package main

import "fmt"

func main() {
	var a = 4
	var ptr *int
	fmt.Printf("[a type is %T]\n",a)
	fmt.Printf("[ptr type is %T]\n",ptr)

	ptr  = &a
	fmt.Printf("[a value is %d]\n",a)
	fmt.Printf("[*ptr value is %d]\n",*ptr)
	fmt.Printf("[ptr value is %v",ptr)
	fmt.Printf("[a pointer address is %p]",&a)

	/*
	%d //整数占位,十进制占位
	%b //显示二进制
	%08d //显示8个占位
	%T //打印类型
	%v //打印原本value
	%p //显示指针位置
	*/

}
