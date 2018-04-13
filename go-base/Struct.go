package main

import "fmt"

type person struct {
	name string
	age  int
}

type skill struct {
	skillname string
}

type Actor struct {
	person
	skill
}

func main() {

	var p person
	p.name = "vzard"
	p.age = 21

	var actor Actor
	actor.person = p
	actor.skill = skill{skillname: "sing"}

	fmt.Println("the person's name is :", p.name)
	fmt.Println("the actor's name is", actor.person.name)
	fmt.Println("the actor's skill is", actor.skill.skillname)

}
