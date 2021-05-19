package main

import "fmt"

type Animal interface {
	Eat()
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (c Cat) Eat() {
	fmt.Printf("%s 正在吃饭\n", c.name)
}

func (c Cat) Sleep() {
	fmt.Printf("%s 正在睡觉\n", c.name)
}

func (d Dog) Eat() {
	fmt.Printf("%s 正在吃饭\n", d.name)
}

func main() {
	var a Animal
	c := Cat{name: "kitty"}
	d := Dog{name: "101"}
	a = c
	a.Eat()
	a = d
	a.Eat()
}
