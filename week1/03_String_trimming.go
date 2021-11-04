package main

import (
	"fmt"
	"strings"
	"unicode"
)
//TrimLeft会去掉连续的cutset
//TrimPrefix会去掉的单一的
func main() {
	x := "@@@@abc_hello_xxx@"
	fmt.Println(strings.Trim(x, "@"))
	fmt.Println(strings.TrimLeft(x, "@"))
	fmt.Println(strings.TrimRight(x, "@"))
	fmt.Println(strings.TrimSpace(x))
	fmt.Println(strings.TrimPrefix(x, "@@@"))
	fmt.Println(strings.TrimSuffix(x, "@"))


	f := func(r rune) bool {
		//if r is Chinese character ,return true
		return unicode.Is(unicode.Han,r)
	}
	fmt.Println(strings.TrimFunc("你好呀，欢迎你的到来",f))
	//x := "@a@ahello_www@"
	fmt.Printf("[TrimLeft :%v]\n",strings.TrimLeft(x,"@a"))
	fmt.Printf("[TrimPrefix :%v]\n",strings.TrimPrefix(x,"@a"))

}
