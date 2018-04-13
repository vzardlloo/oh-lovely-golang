package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

//首字符大写为公有,小写为私有
func (h *Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

//overwrite method
func (e *Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s,Call me on %s", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Japan"}

	mark.SayHi()
	sam.SayHi()
}
