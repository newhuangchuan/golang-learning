package main

import "fmt"

func f5() (s3 []int) {

	s3 = []int{1, 1}
	defer func() {
		s3[1] = 10
	}()
	return []int{3, 3}
}
//改造这个函数
//1. 先给返回值赋值
//2.把defer改造成正常的函数
//3.空洞的return返回

//改造
func f55() (s1 []int) {
	s1 = []int{1, 1}
	s1 = []int{3, 3}
	func() {
		s1[1] = 10 //闭包 s1 [3, 10]
	}()
	return
}

func main() {
	fmt.Println(f5())
	fmt.Println(f55())

}
