package main

import (
	"fmt"
)

func main()  {
	var a int = 10
	var b int = 20
	var c int
	fmt.Printf("please init number is : a = %d, b = %d  c = %d \n",a,b,c)
	//加法
	c = a + b
	fmt.Printf("a + b = %d\n",c)
	//减法
	c = a - b
	fmt.Printf("a - b = %d\n",c)
	//乘法
	c = a * b
	fmt.Printf("a * b = %d\n",c)
	//除法
	c = a / b
	fmt.Printf("a / b = %d\n",c)
	//取余
	c = a % b
	fmt.Printf("a 取余 b = %d\n",c)

	//自增
	c++
	fmt.Printf("[c++ = %d]\n",c)
	//自减
	c--
	fmt.Printf("[c-- = %d]\n",c)
}
