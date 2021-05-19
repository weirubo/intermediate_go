package main

import "fmt"

type Animal interface {
	Action() string
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (c Cat) Action() string {
	return fmt.Sprintf("Cat %s 正在吃饭", c.name)
}

func (d Dog) Action() string {
	return fmt.Sprintf("Dog %s 正在吃饭", d.name)
}

func AnimalAction(a Animal) {
	fmt.Println(a.Action())
}

func main() {
	c := Cat{name: "Kitty"}
	AnimalAction(c)
	d := Dog{name: "101"}
	AnimalAction(d)
}
