package main

import (
	"fmt"
	"net/http"
	"time"
)

// net/http 包
// 编写一个处理器
/*type MyHandler struct {

}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}*/

// 编写多个处理器，处理请求
/*type FirstHandler struct {

}

func (f *FirstHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FirstHandler")
}

type SecondHandler struct {

}

func (s *SecondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SecondHandler")
}*/

// 编写多个函数，处理请求
/*func first(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "func first")
}

func second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "func second")
}*/

// 读取请求首部
// r.Header
// map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9] Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8] Connection:[keep-alive] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36]]
/*func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	h2 := r.Header["User-Agent"]
	fmt.Fprintln(w, h2)
	h3 := r.Header.Get("User-Agent")
	fmt.Fprintln(w, h3)
}*/

// 读取请求主体的数据
// r.ContentLength
// r.Body.Read()
/*func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}*/

// form
// enctype 属性：
// application/x-www-form-urlencoded 默认值，一般用于简单文本数据
// multipart/form-data 一般用于上传文件
// text/plain HTML5 支持

// GET 请求，表单数据以键值对的形式包含在请求的 URL 里面。
// POST 请求，通过主体传递。

// net/http 处理表单数据
/*
使用 Request 结构的方法获取表单数据：
1. 调用 ParseForm 方法或者 ParseMultipartForm 方法，对请求进行语法分析。
2. 取值：
r.Form，map 类型，键是字符串，值是字符串切片。如果键同时存在表单和 URL，值包含表单值和 URL 值，并且表单值排在前面。
r.PostForm，如果键同时存在表单和 URL，只取要表单的值。只支持 application/x-www-form-urlencoded 编码。
r.MultipartForm，支持 multipart/form-data 编码。 只取表单的值，不取 URL 的值。
*/
/*func getVal(w http.ResponseWriter, r *http.Request) {
	// 语法分析
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm)
	fmt.Fprintln(w, r.FormValue("username")) // 只获取第一个值
	fmt.Fprintln(w, r.PostFormValue("username")) // 只获取 form 表单值
}*/

/*func getMultipart(w http.ResponseWriter, r *http.Request) {
	// 语法分析
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm) // 只取表单值，不取 URL 值
	fmt.Fprintln(w, r.MultipartForm)
	fmt.Fprintln(w, r.FormValue("username")) // 只取 URL 值
	fmt.Fprintln(w, r.PostFormValue("username")) // 只取 form 表单值
}*/

/*
FormValue 方法直接获取指定键的值，不需要在之前调用语法分析的方法。如果键同时存在表单和 URL，只取表单的值。
PostFormValue 方法只会取表单的值，不取 URL 的值。
*/

// 文件上传

// JSON 主体
// application/json

// ResponseWriter 写入
/*
Write 方法，
WriteHeader 方法，
Header 方法，
*/

/*func setVal(w http.ResponseWriter, r *http.Request) {
	str := "Hello World!"
	w.WriteHeader(501) // 设置响应返回的状态码，必须在 Write 方法之前调用。
	w.Write([]byte(str)) // 写入响应主体
}*/

/*func setHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302) // WriteHeader 方法执行完，不能再对首部写入，所以要提前对首部写入。
}*/

// cookie

// 写 Cooie
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:  "c1",
		Value: "val1",
	}

	c2 := http.Cookie{
		Name:  "c2",
		Value: "val2",
	}

	c3 := http.Cookie{
		Name:  "c3",
		Value: "val3",
	}
	w.Header().Set("Set-Cookie", c1.String())

	w.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(w, &c3) // 指针类型
}

// 读 Cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header["Cookie"])
	c1, _ := r.Cookie("c1")
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, r.Cookies())
}

// 删 Cookie
func delCookie(w http.ResponseWriter, r *http.Request) {
	c2 := http.Cookie{
		Name:    "c2",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	}
	http.SetCookie(w, &c2)
}

func main() {
	// 构建服务器
	// 简单方式
	// 第二个参数 handler 如果值是 nil，默认值是 DefaultServeMux
	// http.ListenAndServe("", nil)

	// 定义一个 Server 结构体类型的变量，设置服务器更多配置信息。
	/*server := http.Server{
		Addr: ":8080",
		Handler: nil,
	}
	server.ListenAndServe()*/
	// HTTPS
	// server.ListenAndServeTLS("cert.pem", "key.pem")

	// 接收 HTTP 请求
	// 处理器
	/**
	在 Go 语言中，一个处理器就是一个拥有 ServeHTTP 方法的接口，这个 ServeHTTP 方法需要接收两个参数，第一个参数是一个 ResponseWriter 接口，
	第二个参数是一个指向 Request 结构的指针。

	DefaultServeMux 默认多路复用器是多路复用器 ServeMux 结构的一个实例，ServeMux 也拥有 ServeHTTP 方法。
	所以 DefaultServeMux 既是 ServeMux 结构的实例，也是处理器 Handler 结构的实例，因此 DefaultServeMux 不仅是一个多路复用器，还是一个处理器。
	但是 DefaultServeMux 是一个特殊的处理器，它唯一要做的就是根据请求的 URL 将请求重定向到不同的处理器。
	*/

	// 自定义一个处理器，替代 DefaultServeMux。
	/*handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()*/

	/*
		使用自定义的处理器与服务器进行绑定，启动服务器，不管浏览器访问什么地址，服务器返回的都是同样的响应 Hello World!
		这是因为使用自定义的处理器替代了默认多路复用器 DefaultServeMux，服务器不会再通过 URL 匹配来将请求路由至不同的处理器。

		怎么解决？
		使用多个处理器。
	*/

	// 多个处理器，使用 http 包的 Handle 函数绑定到 DefaultServeMux。
	/*
		为了使用多个处理器去处理不同的 URL，我们不再在 Serve 结构的 Handler 字段中指定处理器，而是让服务器使用默认多路复用器 DefaultServeMux，
		然后通过 http.Handle 函数将处理器绑定到 DefaultServeMux。

		http 包的 Handle 函数实际上是 ServeMux 结构的方法，为了操作便利而创建的函数，调用它们等同于调用 DefaultServeMux 的某个方法。例如，调用
		http.Handle，实际上就是在调用 DefaultServeMux 的 Handle 方法。
	*/

	// 自定义 Handler
	/*first := FirstHandler{}
	second := SecondHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// http.Handle 函数
	http.Handle("/first", &first)
	http.Handle("/second", &second)

	server.ListenAndServe()*/

	/*
		以上我们通过使用 http.Handle 函数，将一个创建的处理器绑定到一个 URL 上，实现使用多个处理器处理不同的 URL。

		不过创建多个处理器，代码有些不简洁，有没有更加简洁的方式？
		处理器函数
	*/

	// 处理器函数
	/*
		http.HandleFunc 函数将自定义函数转换成一个处理器 Handler，并将它与 DefaultServeMux 进行绑定，从而简化创建并绑定 Handler 的工作。
	*/
	/*server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// http.HandleFunc
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)

	server.ListenAndServe()*/

	/*
		使用处理器函数和使用处理器，都可以实现根据请求的 URL 将请求重定向到不同的处理器，而且处理器函数比处理器的代码更为简洁。
		但是也不是完全使用处理器函数代替处理器，因为如果代码已经包含了某个接口或某种类型，我们只需为它们添加 ServeHTTP 方法就可以将它们转变为处理器。
	*/

	// HTTP 请求多路复用器 ServeMux 和 DefaultServeMux 是什么关系？
	// DefaultServeMux 是 ServeMux 的默认实例。

	// 使用其他多路复用器
	/*
		ServeMux 无法使用变量实现 URL 模式匹配，使用三方多路复用器 httprouter 包可以实现 URL 模式匹配。
		此外，还有一个非常优秀的三方多路复用器，gorilla/mux
	*/
	// httprouter

	// 处理 HTTP 请求
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// 读取请求首部
	// http.HandleFunc("/headers", headers)

	// 读取请求主体中的数据
	// http.HandleFunc("/body", body)

	// 获取 form 表单的值
	// http.HandleFunc("/getVal", getVal)

	// 获取 multipart
	// http.HandleFunc("/getMultipart", getMultipart)

	// 写入响应主体
	// http.HandleFunc("/setVal", setVal)

	// JSON

	// 对首部写入
	// http.HandleFunc("/setHeader", setHeader)

	// 写 cookie
	http.HandleFunc("/setCookie", setCookie)

	// 读 Cookie
	http.HandleFunc("/getCookie", getCookie)

	// 删 Cookie
	http.HandleFunc("/delCookie", delCookie)
	server.ListenAndServe()

}
