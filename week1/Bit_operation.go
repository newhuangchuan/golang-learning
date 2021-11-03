package main

import "fmt"

func main() {
	var a uint = 60
	var b uint = 13
	var c uint = 0
	fmt.Printf("[a = %d Binary = %08b]\n",a ,a)
	fmt.Printf("[b = %d Binary = %08b]\n",b ,b)

	c = a & b
	fmt.Printf("[ & operate ] [%d = %d  & %d ][%08b = %08b & %08b]\n",c,a,b,c,a,b)

	c = a | b
	fmt.Printf("[ | operate ] [%d = %d  | %d ][%08b = %08b | %08b]\n",c,a,b,c,a,b)

	c = a ^ b
	fmt.Printf("[ ^ operate ] [%d = %d  ^ %d ][%08b = %08b ^ %08b]\n",c,a,b,c,a,b)

	c = a << b
	fmt.Printf("[ << operate ] [%d = %d  << %d ][%08b = %08b <<S %08b]\n",c,a,b,c,a,b)

	c = a >> b
	fmt.Printf("[ >> operate ] [%d = %d  >> %d ][%08b = %08b >> %08b]\n",c,a,b,c,a,b)


}
