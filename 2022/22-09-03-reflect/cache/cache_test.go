package cache

import "testing"

func BenchmarkPopulateStructReflectCache(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructReflectCache(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B(42)", m.B)
		}
	}
}
