package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ch1 := "like golang!!"
	ch2 := "汉字"
	ch3 := "li汉ke字"

	fmt.Printf("[string :%v  bit size or byte number: %d real number :%d]\n",ch1,len(ch1),utf8.RuneCountInString(ch1))
	fmt.Printf("[string :%v  bit size or byte number: %d real number :%d]\n",ch2,len(ch2),utf8.RuneCountInString(ch2))
	fmt.Printf("[string :%v  bit size or byte number: %d real number :%d]\n",ch3,len(ch3),utf8.RuneCountInString(ch3))

}
