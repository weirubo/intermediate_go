package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// log 标准库
// log 包实现了一个简单的 logging 包。
// 它定义了一个 Logger 类型，包含格式化输出的方法。
// 它还有通过帮助函数 Print[f|ln]、Fatal[f|ln]、Panic[f|ln] 访问的预定义标准的 Logger，比手动创建 Logger 更容易使用。
// 该记录器将写入标准错误，并打印每个记录的消息的日期和时间。
// 每个日志消息都在单独的行上输出：如果要打印的消息没有以换行符结尾，则 Logger 将添加一个。
// 写入日志消息后，Fatal 函数将调用 os.Exit(1)。
// Panic 函数会在写入日志消息后调用 panic。
func main () {
	// log.Print("this is a log")
	// log.Printf("error:%s", "this is a error")
	// log.Println("this is a error log")

	// log.Fatal("this is a log")
	// log.Fatalf("error:%s", "this is a error")
	// log.Fatalln("this is a error log")

	// log.Panic("this is a log")
	// log.Panicf("error:%s", "this is a error")
	// log.Panicln("this ia a error log")

	logFile, err := os.OpenFile("error1.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()
	// log.SetOutput(logFile)
	// log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	// log.SetFlags(log.LstdFlags|log.Lshortfile)
	// log.SetPrefix("[DEBUG] ")
	// log.Println("this is a test log of write in file")
	logs := DefinesLogger(logFile, "", log.LstdFlags|log.Lshortfile)
	logs.Debug("message")
	logs.Debugf("%s", "content")
}

// 自定义 logger
type Logger struct {
	definesLogger *log.Logger
}

type Level int8

const(
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return " [debug] "
	case LevelInfo:
		return " [info] "
	case LevelError:
		return " [error] "
	}
	return ""
}

func DefinesLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{definesLogger: l}
}

func (l *Logger) Debug(v ...interface{}) {
	l.definesLogger.Print(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.definesLogger.Print(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.definesLogger.Print(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.definesLogger.Print(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.definesLogger.Print(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.definesLogger.Print(LevelError, fmt.Sprintf(format, v...))
}