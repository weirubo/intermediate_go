package main

import "fmt"

// 指针

func main() {
	/**
	指针是一种数据类型，用来存储值的内存地址，为了便于理解，我们也可以把指针理解为内存地址，指针类型只占用内存 4 个或 8 个字节，在 Golang 语言中，类型名称前加 `*` 表示对应类型的指针类型。

	指针类型变量也需要一块内存空间存储值，指针变量的值就是它所指向数据的内存地址，普通变量的值就是具体存放的数据。不同的指针类型变量之间无法互相赋值，

	在 Golang 语言中，指针不支持运算，也不能获取常量的指针。
	*/

	// 方式 1
	// 使用取地址符 & 获取变量的指针（内存地址）。

	// 定义普通变量 a
	a := 1
	// 定义指针变量 p
	p := &a
	fmt.Println("变量 a 的值为：", a)     // 1
	fmt.Println("变量 a 的内存地址为：", p)  // 0xc0000ae008
	fmt.Printf("变量 a 的类型为：%T\n", a) // int
	fmt.Printf("变量 p 的类型为：%T\n", p) // *int

	// 方式 2

	// 使用 var 关键字声明指针变量，使用 var 关键字声明的变量不能直接赋值和取值，因为它还没有内存地址，它的值是 nil。
	// var str string
	// var p1 *int
	// 不同指针类型变量之间无法互相赋值
	// p1 = &str // ./main.go:29:5: cannot use &str (type *string) as type *int in assignment

	// 方式 3

	// 使用内置的 new 函数来声明指针类型的变量，new 函数接收一个参数，可以传递类型给它，返回值是传递类型的指针类型。
	p2 := new(int)
	fmt.Printf("%v %T\n", p2, p2)

	// ==========

	// 指针操作 - 取值，修改

	// 获取指针指向的值
	// 想要获取指针指向的值，只需在指针变量前加 `*`。
	b := 2
	p3 := &b
	val := *p3
	fmt.Println("变量 val 的值为：", val)

	// 修改指针指向的值
	// 给 *p3 赋值，*p3 指向的值也被修改，因为 p3 指向的内存就是变量 b 的内存地址。
	*p3 = 3
	fmt.Println("*p3 指针指向的值为：", *p3)
	fmt.Println("变量 b 的值为：", b)

	// 使用 var 关键字声明的指针变量不能直接赋值和取值，因为它还没有分配内存，它的值为 nil，可以使用内置函数 new 给它分配内存。
	var p4 *int = new(int)
	*p4 = 4
	fmt.Println(*p4)

	// ==========

	// 指针参数
	/**
	在 Golang 语言中，函数传递参数只有值传递，传递的都是参数的拷贝副本，所以我们传递值类型的参数时，修改参数的值，原始数据不会被修改。但是，如果是指针类型的参数，修改参数的值，原始数据也会被修改，原因是指针类型的参数存储的是内存地址。
	*/

	// 值类型参数，实参的值未改变
	mySalary := 80000
	fmt.Printf("变量 mySalary 的内存地址为：%p\n", &mySalary)
	modifySalary(mySalary)
	fmt.Println(mySalary)

	// 指针类型参数，实参的值被改变
	modifySalary2(&mySalary)
	fmt.Println(mySalary)

	// why
	// 因为指针类型参数的内存地址和实参的内存地址相同。

	// ==========

	// 指针接收者

	/**
	如果需要修改接收者，可以使用指针修改指针指向数据的值。

	如果接收者是非 map、slice 和 channel 类型，并且数据比较大，可以使用指针来节省内存。
	*/

	// 值类型调用者
	w := worker{
		name:   "frank",
		salary: 5000,
	}
	// 指针类型接收者
	w.raise()
	fmt.Printf("w 的姓名是 %s，薪水是每月 %d\n", w.name, w.salary)

	// 值类型调用者
	w1 := worker{
		name:   "frank1",
		salary: 5000,
	}
	// 值类型接收者
	w1.raise1()
	fmt.Printf("w1 的姓名是 %s，薪水是每月 %d\n", w1.name, w1.salary)

	// 指针类型调用者
	w2 := &worker{
		name:   "lucy",
		salary: 5000,
	}
	// 指针类型接收者
	w2.raise()
	fmt.Printf("w2 的姓名是 %s，薪水是每月 %d\n", w2.name, w2.salary)

	// 指针类型调用者
	w3 := &worker{
		name:   "lucy1",
		salary: 5000,
	}
	// 值类型接收者
	w3.raise1()
	fmt.Printf("w3 的姓名是 %s，薪水是每月 %d\n", w3.name, w3.salary)
}

func modifySalary(salary int) {
	fmt.Printf("参数变量的内存地址为：%p\n", &salary)
	salary = 100000
}

func modifySalary2(salary *int) {
	fmt.Printf("参数变量的内存地址为：%p\n", salary)
	*salary = 100000
}

type worker struct {
	name   string
	salary uint
}

func (w *worker) raise() {
	w.salary += 1000
}

func (w worker) raise1() {
	w.salary += 1000
}
