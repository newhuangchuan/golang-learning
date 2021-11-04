package main

import "fmt"

func main() {

	ch3 := "li汉ke字"

	//Subscript traversal
	for i := 0; i < len(ch3) ; i++ {
		fmt.Printf("[ascii: %c  %d]\n",ch3[i],ch3[i])
	}

	fmt.Println("=====================================")
	//Traverse
	for _, i := range ch3{

		fmt.Printf("[unicode: %c %d]\n",i, i)
	}

}
