package unsafe

import "testing"

func BenchmarkPopulateStructReflectCache(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B(42)", m.B)
		}
	}
}

func BenchmarkPopulateStructReflectCache2(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe2(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B(42)", m.B)
		}
	}
}

func BenchmarkPopulateUnsafe3(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	descriptor, err := describeType((*SimpleStruct)(nil))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe3(&m, descriptor); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B(42)", m.B)
		}
	}
}

//  BenchmarkPopulateUnsafe3 almost equal BenchmarkPopulate, amazing !!!!
func BenchmarkPopulate(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		populateStruct(&m)
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B(42)", m.B)
		}
	}
}
