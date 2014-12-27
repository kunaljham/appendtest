package appendtest

import "testing"

func BenchmarkAppend1(b *testing.B) {
	ifs := make([]interface{}, b.N)
	ss := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		appendStringsToInterfaceSlice1(ifs, ss)
	}
	b.ReportAllocs()
}

func BenchmarkAppend2(b *testing.B) {
	ifs := make([]interface{}, b.N)
	ss := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		appendStringsToInterfaceSlice2(ifs, ss)
	}
	b.ReportAllocs()
}

func BenchmarkAppend3(b *testing.B) {
	ifs := make([]interface{}, b.N)
	ss := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		appendStringsToInterfaceSlice3(ifs, ss)
	}
	b.ReportAllocs()
}
