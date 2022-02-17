package main

import "fmt"

func main() {
	s1 := make([]int, 0)
	s1 = append(s1, 100)
	s1 = append(s1, 101)
	s1 = append(s1, 103)
	s1 = append(s1, 104)

	fmt.Println(s1)

	//init len == 5 , cap == 5 slice

	s2 := make([]int, 5, 5)
	s2 = append(s2, 111)
	s2 = append(s2, 114)
	s2 = append(s2, 116)
	s2 = append(s2, 117)

	fmt.Println(s2)
	fmt.Printf("[cap : %d , len : %d ]\n", cap(s2), len(s2))
}
