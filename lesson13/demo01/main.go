package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Eat() {
	fmt.Printf("%s 正在吃饭\n", c.name)
}

func main() {
	c := Cat{name: "kitty"}
	c.Eat()
}
