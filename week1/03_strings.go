package main

import (
	"fmt"
	"strings"
)

func main() {
	//Determine whether there is a character or substring

	//Find substring
	fmt.Println(strings.Contains("localhost:8080","8080"))
	//Any character
	fmt.Println(strings.Contains("localhost:8080","a b"))
	//Find rune
	fmt.Println(strings.Contains("北京欢迎你","你"))




}
