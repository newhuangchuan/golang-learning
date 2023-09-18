package main

import "log"

type Person struct {
	Name string
	Age int
	Labels map[string]string
}

func main() {
	var p Person
	log.Printf("%v", p)
	log.Printf("%+v", p)

	p1 := Person{
		Name:   "admin",
		Age:    1999,
		Labels: map[string]string{
			"kubernetes.io":"helm",
		},
	}
	log.Printf("%+v", p1)
}