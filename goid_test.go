package goid_test

import (
	"testing"

	"github.com/aauutthh/goid"
)

func BenchmarkGetGoIDASM(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		goid.GetGoIDAsm()
	}
}

func BenchmarkGetGoID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		goid.GetGoID()
	}
}
