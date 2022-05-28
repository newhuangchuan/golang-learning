package main

import "fmt"

func main() {
	defer function(1,function(3, 0))
	defer function(1,function(4, 0))

}

func function(index int , value int) int {
	fmt.Println(index)
	return index
}
//先defer函数压栈，第一个defer
// 先去计算函数的参数 3 然后压栈
// 执行函数，得到4作为参数压栈 3 4
