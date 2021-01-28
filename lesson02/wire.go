//+build wireinject

package main

import "github.com/google/wire"

// 注入器
// 一个应用程序使用注入器将这些提供程序连接起来：以依赖关系顺序调用提供程序的函数。
// 只要返回值的类型正确，返回值就无关紧要。这些值本身将在生成的代码中被忽略。
// 注入程序可以在输入上进行参数化（然后将其发送给提供程序），并且可以返回错误。
// wire.Build 的参数与 wire.NewSet 相同：它们形成提供程序集。这是在生成该注入器的代码期间使用的提供者集。
func InitializeEvent(phrase string) (Event, error) {
	// 使用 Wire，编写注入器的签名，然后 Wire 生成函数的主体。
	// 通过编写函数声明来声明注入器，该函数声明的主体是对 wire.Build 的调用。
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}

// 在带有注入器的文件中找到的所有非注入器声明都将被复制到生成的文件中。
// 您可以通过在软件包目录中调用 Wire 来生成注入器
// Wire 将在名为 wire_gen.go 的文件中生成注入器的实现，输出非常接近开发人员自己编写的内容。
// 此外，在运行时对 Wire 的依赖性很小：所有编写的代码只是普通的 Go 代码，无需 Wire 即可使用。
// 创建 wire_gen.go 后，您可以通过运行 go generate 重新生成它。