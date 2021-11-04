package main

import (
	"fmt"
	"strings"
)

func main() {

	baseUrl := "http://localhost:8080/api/v1/query?"

	args := strings.Join([]string{"name=root","id=1","env=online"},"&")
	fulluri := baseUrl + args
	fmt.Println(fulluri)

}
