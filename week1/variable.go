package main

import "fmt"

const (

	_ = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
)

func T1() string {
	str1 := "function T1 string variable"
	fmt.Println(str1)
	return str1
}

func main() {

	fmt.Println(KB,MB,GB,TB)
	fmt.Println(1 << 10)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	if str := T1();str == "" {
		fmt.Println("function return null")
	}


}
