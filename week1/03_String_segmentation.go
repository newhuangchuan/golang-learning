package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "inf.bigdata.kafka"

	s2 := "localhost:8080/api/v1/host/1"
	ss1 := strings.Split(s1,".")
	ss2 := strings.SplitAfter(s1,".")


	ps1 := strings.Split(s2,"/")
	ps2 := strings.SplitAfterN(s2,"/",2)

	fmt.Printf("[split :%v ]\n",ss1)
	fmt.Printf("[splitafter :%v ]\n",ss2)
	fmt.Printf("[split uri :%v ]\n",ps1)
	fmt.Printf("[splitaftern uri :%v ]\n",ps2)

}
