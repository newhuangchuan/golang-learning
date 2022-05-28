package main

import "log"

type Person2 struct {
	Name string
	Age int
	Label map[string]string
}

type Student1 struct {
	StudentId int
	Person2
}

type  Teacher struct {
	TeacherId int
	P Person2
}
func main() {
	 p1 := Person2{
	 	Name: "tty",
	 	Age: 99,
	 	Label: map[string]string{},
	 }
	//匿名结构体的嵌入
	 s1 := Student1{
	 	StudentId: 994399223,
	 	Person2: p1,
	 }

	 p2 := Person2{
	 	Name: "tty2",
	 	Age: 89,
	 	Label: map[string]string{},
	 }
	//命名结构体的嵌入
	 t1 := Teacher{
	 	TeacherId: 2287334443,
	 	P: p2,
	 }
	//匿名结构体可以直接访问继承的属性
	 log.Printf("[s1.name is %v,s1.age is %v, s1.studentid is %v]\n", s1.Name, s1.Age, s1.StudentId)
	 log.Printf("[s1.Person2.name is %v, s1.Person2.age is %v, s1.studentid is %v]\n", s1.Person2.Name, s1.Person2.Age,s1.StudentId)
	 log.Printf("[t1.P.name is %v,t1.P.age is %v, t1.studentid is %v]\n",t1.P.Name,t1.P.Age,t1.TeacherId )
}
