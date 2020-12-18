package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Wire 是一种代码生成工具，可以使用依赖注入自动连接组件。
// 组件之间的依赖关系在 Wire 中表示为函数参数，鼓励显式初始化而不是全局变量。
// 因为 Wire 在运行时没有运行时状态或反射，所以编写用于 Wire 的代码甚至对于手写初始化也是有用的。

// 安装 Wire
// go get github.com/google/wire/cmd/wire
// 并确保将 $GOPATH/bin 添加到您的 $PATH中。

type Message string

// 初始化
// 提供者 provider 一个有返回值的函数
func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter  struct {
	Message Message
	Grumpy bool
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// 初始化
// 提供者 provider
// 一个有返回值且可导出的函数，
// 通过使用参数指定依赖项，返回一个 Greeter，依赖一个 Message
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix() % 2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

type Event struct {
	Greeter Greeter
}

// 初始化
// 一个有返回值且可导出的函数，
// 通过使用参数指定依赖项，返回一个 Greeter，依赖一个 Message，
// 提供者也可以返回错误，
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

// 可以将经常一起使用的提供者分组为提供者集合，
// 使用 wire.NewSet 函数，将这些提供者添加到名为 SuperSet 的新集合中，并且可以在一个提供者集合中嵌套其它的提供者集合，

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main () {
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)
	// event.Start()

	e, err := InitializeEvent("testStr")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}