package main

import "fmt"

func main() {

	var arr1 [5]string
	fmt.Println(arr1)

	arr1[1] = "golang"
	fmt.Println(arr1)


	var arr2 = [...]int{1,2,3}
	fmt.Println(arr2)

}
