package string_test

import (
	"bytes"
	"strings"
	"testing"
)

var (
	concatSteps = 1000
	subStr      = "s"
	expectedStr = strings.Repeat(subStr, concatSteps)
)

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var s string
		for i := 0; i < concatSteps; i++ {
			s += subStr
		}
		if s != expectedStr {
			b.Errorf("unexpected result, got: %s, want: %s", s, expectedStr)
		}
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buffer bytes.Buffer
		for i := 0; i < concatSteps; i++ {
			buffer.WriteString(subStr)
		}
		if buffer.String() != expectedStr {
			b.Errorf("unexpected result, got: %s, want: %s", buffer.String(), expectedStr)
		}
	}
}

func BenchmarkBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var builder strings.Builder
		for i := 0; i < concatSteps; i++ {
			builder.WriteString(subStr)
		}
		if builder.String() != expectedStr {
			b.Errorf("unexcepted result, got: %s, want: %s", builder.String(), expectedStr)
		}
	}
}

func BenchmarkCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bytes := make([]byte, len(subStr)*concatSteps)
		c := 0
		for i := 0; i < concatSteps; i++ {
			c += copy(bytes[c:], subStr)
		}
		if string(bytes) != expectedStr {
			b.Errorf("unexpected result, got: %s, want: %s", string(bytes), expectedStr)
		}
	}
}
