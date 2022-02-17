package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)
	s1 = append(s1, 10)
	s1 = append(s1, 11)
	s1 = append(s1, 12)
	fmt.Println(s1)

	var s2 = []int{1, 2, 3}
	fmt.Println(s2)

}
