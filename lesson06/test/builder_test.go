package test

import (
	"strings"
	"testing"
)

var lan []string = []string{
	"golang",
	"php",
	"javascript",
}

func stringBuilder(lan []string) string {
	var str strings.Builder
	for _, val := range lan {
		str.WriteString(val)
	}
	return str.String()
}

func stringBuilderGrow(lan []string) string {
	var str strings.Builder
	str.Grow(16)
	for _, val := range lan {
		str.WriteString(val)
	}
	return str.String()
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilder(lan)
	}
}

func BenchmarkBuilderGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilderGrow(lan)
	}
}
