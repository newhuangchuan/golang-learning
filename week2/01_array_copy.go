package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 [3]int
	arr2 = arr1

	fmt.Printf("[[%v  %p]]\n", arr1, &arr1)
	fmt.Printf("[[%v  %p]]\n", arr2, &arr2)

	arr2[2] = 20
	fmt.Printf("[[%v  %p]]\n", arr1, &arr1)
	fmt.Printf("[[%v  %p]]\n", arr2, &arr2)

}
