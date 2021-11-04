package main

import "fmt"

func main() {

	s2 := "localhost:8080"
	fmt.Println(s2)
	//Forced type conversion   string to []byte
	sByte := []byte(s2)
	//subscript modification
	sByte[len(sByte) - 1] = '1'

	//Forced type conversion    []byte to string

	s3 := string(sByte)
	fmt.Println(s3)

}
