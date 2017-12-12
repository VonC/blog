package main

import (
	"testing"
)

func Benchmark_createImage(b *testing.B) {
	c := complex(0.312, 0.5)
	for i := 0; i < b.N; i++ {
		createImage(size, limit, c)
	}
}
