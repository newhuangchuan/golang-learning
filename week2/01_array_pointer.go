package main

import (
	"fmt"
)

func main() {

	var arr1 [5]*int
	//根据索引赋值
	arr1[0] = new(int)
	arr1[1] = new(int)
	arr1[2] = new(int)
	arr1[3] = new(int)
	arr1[4] = new(int)
	fmt.Println(arr1)
	*arr1[0] = 10
	*arr1[1] = 2
	fmt.Println(arr1)

	for i := 0 ;i < len(arr1); i++ {
		fmt.Println("index : " , i , "value : ", *arr1[i])
	}

}
