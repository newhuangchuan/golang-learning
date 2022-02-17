package main

import "fmt"

func showSlice(s []int) {
	fmt.Printf("[send slice is : %v]\n", s)
	s[2] = 40
}

func main() {
	s := []int{1, 3, 4, 6}

	showSlice(s)

	fmt.Println(s)
	//虽然函数传参是指传递，应该是深拷贝
	//但是slice属于引用类型，浅拷贝，在函数内部的修改还是会影响外部
}
