package main

import "fmt"

// 封装

type Employee struct {
	Name   string
	Gender string
	Age    uint8
	Salary uint
}

// 方法

func (e Employee) Information() string {
	return fmt.Sprintf("Name:%s Gender:%s Age:%d Salary:%d", e.Name, e.Gender, e.Age, e.Salary)
}

func (e *Employee) InformationPointer() string {
	return fmt.Sprintf("Name:%s Gender:%s Age:%d Salary:%d", e.Name, e.Gender, e.Age, e.Salary)
}

func (e Employee) SalaryIncr() uint {
	e.Salary = e.Salary + 1000
	return e.Salary
}

func (e *Employee) SalaryIncrPointer() uint {
	e.Salary = e.Salary + 1000
	return e.Salary
}

func main() {
	// 实例化
	e1 := Employee{
		Name:   "lucy",
		Gender: "lady",
		Age:    28,
		Salary: 5000,
	}
	e1Information := e1.Information()
	fmt.Println(e1Information)

	e1InformationPointer := e1.InformationPointer()
	fmt.Println(e1InformationPointer)

	e1SalaryIncr := e1.SalaryIncr()
	fmt.Println(e1SalaryIncr, e1)

	e1SalaryIncrPointer := e1.SalaryIncrPointer()
	fmt.Println(e1SalaryIncrPointer, e1)
}
