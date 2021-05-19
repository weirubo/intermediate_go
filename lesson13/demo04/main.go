package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Cat struct{}

func (c Cat) Eat() {
	fmt.Println("Cat 正在吃饭")
}

func (c Cat) Sleep() {
	fmt.Println("Cat 正在睡觉")
}

type Dog struct{}

func (d Dog) Eat() {
	fmt.Println("Dog 正在吃饭")
}

func (d Dog) Sleep() {
	fmt.Println("Dog 正在睡觉")
}

func main() {
	var _ Animal = (*Cat)(nil)
	var _ Animal = (*Dog)(nil)
	c := Cat{}
	c.Eat()
	d := Dog{}
	d.Eat()
}
