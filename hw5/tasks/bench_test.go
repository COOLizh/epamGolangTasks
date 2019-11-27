package tasks

import (
	"testing"
)

func Benchmark_MyStrToInt(b *testing.B) {
	MyStrToInt("-9223372036854775808")
}

func Benchmark_MyStrToInt2(b *testing.B) {
	MyStrToInt2("-9223372036854775808")
}
