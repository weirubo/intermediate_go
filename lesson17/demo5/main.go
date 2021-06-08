package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("text.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	n, err := f.WriteString("this is a text file\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
