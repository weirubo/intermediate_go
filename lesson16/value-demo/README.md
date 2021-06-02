>「[Golang 语言 method 接收者使用值类型和指针类型的区别](https://mp.weixin.qq.com/s/OOCiJlCOKe6EE98nydBExw)」Part 04 错误修正

######值类型怎么避免拷贝副本
阅读到这里，读者朋友可能会简单认为使用值类型会拷贝副本，使用指针类型不会拷贝副本。实际上，我们可以通过优化代码，在不改变语义的前提下，实现使用值类型也不会拷贝副本。

示例代码 v1：
```golang
type User struct {
	name string
	nickName string
}

func (u *User) SetName (str string) {
	u.name = str
}

func (u *User) SetNickName (str string) {
	u.nickName = str
}

func (u *User) PrintUserInfo () {
	fmt.Printf("name:%s nickName:%s\n", u.name, u.nickName)
}

func main () {
    user := &User{}
    fmt.Printf("user pointer:%p\n", user)
    for i := 0; i < 10; i++ {
        user.SetName(fmt.Sprintf("name-%d", i))
        user.SetNickName(fmt.Sprintf("nickName-%d", i))
        user.PrintUserInfo()
    }
}
```
output:
```golang
user pointer:0xc0000b6000
name:name-0 nickName:nickName-1
name:name-1 nickName:nickName-2
name:name-2 nickName:nickName-3
name:name-3 nickName:nickName-4
name:name-4 nickName:nickName-5
name:name-5 nickName:nickName-6
name:name-6 nickName:nickName-7
name:name-7 nickName:nickName-8
name:name-8 nickName:nickName-9
name:name-9 nickName:nickName-10
```
阅读上面这段代码，我们使用指针类型的接收者定义了两个方法，for-loop 循环执行 10 次。输出结果也是我们想要的。

示例代码 v2：
```golang
type User struct {
	name string
	nickName string
}

func (u *User) SetName (str string) {
	u.name = str
}

func (u *User) SetNickName (str string) {
	u.nickName = str
}

func (u *User) PrintUserInfo () {
	fmt.Printf("name:%s nickName:%s\n", u.name, u.nickName)
}

func main () {
    user := &User{}
    fmt.Printf("user pointer:%p\n", user)
    for i := 0; i < 10; i++ {
        go func(i int) {
        user.SetName(fmt.Sprintf("name-%d", i))
        user.SetNickName(fmt.Sprintf("nickName-%d", i + 1))
        user.PrintUserInfo()
        }(i)
    }
    time.Sleep(1 * time.Second)
}
```
output：
```golang
user pointer:0xc000068020
name:name-9 nickName:nickName-10
name:name-7 nickName:nickName-8
name:name-8 nickName:nickName-9
name:name-5 nickName:nickName-6
name:name-0 nickName:nickName-1
name:name-2 nickName:nickName-3
name:name-3 nickName:nickName-4
name:name-4 nickName:nickName-7
name:name-6 nickName:nickName-7
name:name-1 nickName:nickName-2

```
阅读上面这段代码，我们发现多次运行应用程序，偶尔会出现错误的输出结果。

示例代码 v3：
```golang
type User struct {
	name string
	nickName string
}

func (u *User) SetName (str string) {
	u.name = str
}

func (u *User) SetNickName (str string) {
	u.nickName = str
}

func (u *User) PrintUserInfo () {
	fmt.Printf("name:%s nickName:%s\n", u.name, u.nickName)
}

func main () {
    for i := 0; i < 10; i++ {
        user := &User{}
        fmt.Printf("user-%d pointer:%p\n",i, user)
        go func(i int) {
            user.SetName(fmt.Sprintf("name-%d", i))
            user.SetNickName(fmt.Sprintf("nickName-%d", i + 1))
            user.PrintUserInfo()
        }(i)
    }
    time.Sleep(1 * time.Second)
}
```
output:
```golang
user-0 pointer:0xc000068020
user-1 pointer:0xc000068040
user-2 pointer:0xc000068060
user-3 pointer:0xc000068080
user-4 pointer:0xc0000680a0
user-5 pointer:0xc0000680c0
user-6 pointer:0xc0000680e0
user-7 pointer:0xc000068100
user-8 pointer:0xc000068120
user-9 pointer:0xc000068140
name:name-9 nickName:nickName-10
name:name-6 nickName:nickName-7
name:name-7 nickName:nickName-8
name:name-8 nickName:nickName-9
name:name-0 nickName:nickName-1
name:name-2 nickName:nickName-3
name:name-4 nickName:nickName-5
name:name-1 nickName:nickName-2
name:name-5 nickName:nickName-6
name:name-3 nickName:nickName-4

```
阅读上面这段代码，我们将实例化 User 放在 for-loop 中，这样每次循环都会重新实例化 User（拷贝副本），虽然也是可以得到我们想要的输出结果。但是每次循环都会实例化 User（拷贝副本），从而会浪费内存空间。

示例代码 v4：
```golang
type User struct {
	name string
	nickName string
}
func (u User) SetName (str string) User {
    u.name = str
    return u
}

func (u User) SetNickName (str string) User {
    u.nickName = str
    return u
}

func (u User) PrintUserInfo () {
    fmt.Printf("name:%s nickName:%s\n", u.name, u.nickName)
}

func main () {
    user := &User{}
    fmt.Printf("user pointer:%p\n", user)
    for i := 0; i < 10; i++ {
        go func(i int) {
            user.SetName(fmt.Sprintf("name-%d", i)).SetNickName(fmt.Sprintf("nickName-%d", i + 1)).PrintUserInfo()
        }(i)
    }
    time.Sleep(1 * time.Second)
}
```
output:
```golang
user pointer:0xc000068020
name:name-9 nickName:nickName-10
name:name-7 nickName:nickName-8
name:name-8 nickName:nickName-9
name:name-5 nickName:nickName-6
name:name-2 nickName:nickName-3
name:name-0 nickName:nickName-1
name:name-3 nickName:nickName-4
name:name-4 nickName:nickName-5
name:name-1 nickName:nickName-2
name:name-6 nickName:nickName-7
```
阅读上面这段代码，我们发现输出结果也是我们想要的，并且也不会拷贝副本。原因是我们使用值类型的接收者，并且使 method 返回该接收者，从而做到多个 method 使用同一个接收者，将字段串联在一起，避免了拷贝副本。
