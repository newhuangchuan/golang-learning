package main

import (
	"log"
)

type Person struct {
	Name string
	Age int
	Label map[string]string
}

func main() {
	var p = Person{
		Name: "root",
		Age: 18,
		Label: map[string]string{},
	}
	var p1 *Person = new(Person)
	p1.Age = 19
	p1.Name = "xiao"
	log.Println(*p1)
	p2 := Person{
		Name: "administrator",
		Age: 99,
		Label: map[string]string{},
	}
	p3 := new(Person)
	p3 = &Person{"admin", 18, map[string]string{}}
	log.Printf("[%v  %+v]\n",p, p)
	log.Printf("[%+v]\n", p2)
	log.Printf("[%+v]\n", p3)
}
