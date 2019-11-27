package tasks

import (
	"testing"
)

func Benchmark_MyStrToInt(b *testing.B) {
	str := "-9223372036854775808"
	for i := 0; i < b.N; i++ {
		MyStrToInt(str)
	}
}

func Benchmark_MyStrToInt2(b *testing.B) {
	str := "-9223372036854775808"
	for i := 0; i < b.N; i++ {
		MyStrToInt2(str)
	}
}
