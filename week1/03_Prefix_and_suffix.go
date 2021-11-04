package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "inf.bigdata.kafka"

	fmt.Println(strings.HasPrefix(s1,"inf"))
	fmt.Println(strings.HasSuffix(s1,"kafka"))
	fmt.Println(strings.HasSuffix(s1,""))

}
