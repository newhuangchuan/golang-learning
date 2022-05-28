package main

import "fmt"

type Person3 struct {
	Name string
	Age int
	Label map[string]string
}

type Student2 struct {
	StudentId int
	Person3
}

type  Teacher1 struct {
	TeacherId int
	*Person3  //结构体匿名指针嵌入
}


func main() {
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal 0xc0000005 code=0x0 addr=0x0 pc=0x857a0a]

	*/
	t1 := Teacher1{}
	fmt.Println(t1)

}
