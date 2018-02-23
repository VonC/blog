package main

import (
	"testing"
)

func Benchmark_createImageSimple(b *testing.B) {
	c := complex(0.312, 0.5)
	flagfill = true
	for i := 0; i < b.N; i++ {
		createImage(size, limit, c)
	}
	flagfill = false
}

func Benchmark_createImageGoPerPixel(b *testing.B) {
	c := complex(0.312, 0.5)
	flagfillgopixel = true
	for i := 0; i < b.N; i++ {
		createImage(size, limit, c)
	}
	flagfillgopixel = false
}

func Benchmark_createImageGoPerCol(b *testing.B) {
	c := complex(0.312, 0.5)
	flagfillgorow = true
	for i := 0; i < b.N; i++ {
		createImage(size, limit, c)
	}
	flagfillgorow = false
}
