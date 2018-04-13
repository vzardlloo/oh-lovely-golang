package main

import "fmt"

func main() {

	var boy, girl string = "I'm a boy", "I'm a girl"
	fmt.Println(boy, girl)

	boy1, girl1 := "I'm a boy", "I'm a girl";
	fmt.Println(boy1, girl1)

	_, girl2 := "hahaha", "hehehe"
	fmt.Println(girl2)

	const Pi = 3.14159
	fmt.Println(Pi)
	var enabled = true
	fmt.Println(enabled)

}
