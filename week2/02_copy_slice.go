package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := make([]int, 3)
	a3 := make([]int, 1)

	copy(a2, a1)
	copy(a3, a1)

	a2[0] = 10
	fmt.Println(a1, a2)
	a1[2] = 30
	fmt.Println(a1, a2)
	fmt.Println(a3)
}
