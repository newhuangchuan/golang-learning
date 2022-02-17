package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 5, 6, 7}

	for k, v := range s1 {
		fmt.Printf("[key is %d, value is : %d]\n", k, v)
	}

	for index := range s1 {
		s1[index] += 100
	}

	for index, num := range s1 {
		fmt.Printf("[new key is %d , new value is %d]\n", index, num)
	}
}
