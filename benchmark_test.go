package main

import (
	"github.com/travertischio/addison-wesley/popcount"
	"testing"
)

var count int

func BenchmarkPopCountSingle(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		count = popcount.PopCountSingle(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		count = popcount.PopCountLoop(uint64(i))
	}
}
