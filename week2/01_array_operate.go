package main

import "fmt"

func main() {

	var arr1 [10]int

	//根据索引赋值
	for i := 0; i < 10; i++ {
		arr1[i] = i
	}

	for i := 0; i < 10; i++ {
		fmt.Println(arr1[i])
	}

}
