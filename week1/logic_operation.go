package main

import "fmt"

func main() {

	var a = true
	var b = false

	if a && b {
		fmt.Println("[A and b are both true to be true ]")
	}else {
		fmt.Println("[a or b is not true]")
	}

	if a || b {
		fmt.Println("[A and b have a true is true]")
	}
	if !(a && b) {
		fmt.Println("[a and b are both true, and then the negative is true]")
	}

}
