package main

import (
	"testing"
)

func BenchmarkLogic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logicCode()
	}
}
