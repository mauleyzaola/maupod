package helpers

import "testing"

func generateUUID(n int) {
	keys := make(map[string]struct{})
	for i := 0; i < n; i++ {
		val := NewUUID()
		if _, ok := keys[val]; ok {
			panic("duplicated uuid entry")
		}
		keys[val] = struct{}{}
	}
}

func benchmarkNewUUID(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateUUID(i)
	}
}

func BenchmarkNewUUID100(b *testing.B) {
	benchmarkNewUUID(100, b)
}

func BenchmarkNewUUID500(b *testing.B) {
	benchmarkNewUUID(500, b)
}

func BenchmarkNewUUID1000(b *testing.B) {
	benchmarkNewUUID(1000, b)
}
