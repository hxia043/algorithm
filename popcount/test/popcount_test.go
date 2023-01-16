package main

import (
	"popcount/iterated"
	"popcount/lookup"
	"popcount/sparse"
	"testing"
)

var count uint64

func benchmark(b *testing.B, f func(uint64, uint64)) {
	for i := 0; i < b.N; i++ {
		f(uint64(i), count)
	}
}

func BenchmarkIterated(b *testing.B) { benchmark(b, iterated.Iterated) }

func BenchmarkSparse(b *testing.B) { benchmark(b, sparse.Sparse) }

func BenchmarkLookup(b *testing.B) { benchmark(b, lookup.Lookup) }
