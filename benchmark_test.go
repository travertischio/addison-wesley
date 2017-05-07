package main

import (
	"fmt"
	"github.com/travertischio/addison-wesley/popcount"
	"os"
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

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		count = popcount.PopCountShift(uint64(i))
	}
}

func TestDoTheseWorkSingle(t *testing.T) {
	test := popcount.PopCountSingle(uint64(29))
	if 4 != test {
		fmt.Printf("%v", test)
		os.Exit(1)
	}
}

func TestDoTheseWorkLoop(t *testing.T) {
	test := popcount.PopCountLoop(uint64(29))
	if 4 != test {
		fmt.Printf("%v", test)
		os.Exit(1)
	}
}

func TestDoTheseWorkShift(t *testing.T) {
	test := popcount.PopCountShift(uint64(29))
	if 4 != test {
		fmt.Printf("%v", test)
		os.Exit(1)
	}
}
