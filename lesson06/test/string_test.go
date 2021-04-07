package test

import "testing"

var strs string = `Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.`

func str(str string) {
	_ = str + "golang"
}

func ptr(str *string) {
	_ = *str + "golang"
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str(strs)
	}
}

func BenchmarkStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptr(&strs)
	}
}
