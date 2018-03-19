package main

import (
	"testing"
)

func BenchmarkRequestNats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RequestNats()
	}
}
func BenchmarkRequestGrpc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RequestGrpc()
	}
}
