package main

import (
	"log"
)

type Person1 struct {
	Name string
	Age  int
	Label map[string]string
}

type Student struct {
	StudentId int
	Person1  //匿名结构体的嵌套
}

func main() {
	p1 := Person1{
		Name: "abc",
		Age: 18,
		Label: nil,
	}
	//访问属性
	log.Printf("[p.name : %v][p.age : %v]\n", p1.Name, p1.Age)
	//修改属性
	p1.Age +=1
	log.Printf("[p.name : %v][p.age : %v]\n", p1.Name, p1.Age)

	s1  := Student{
		StudentId: 2292929,
		Person1: p1,
	}

	log.Printf("[s1.name : %v][s1.age : %v]\n", s1.Name, s1.Age)
}
