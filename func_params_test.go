package samples

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkParameterInStruct(b *testing.B) {
	a := A{
		x: "a",
		y: 1,
	}

	ns := make([]N, 100)
	for i := 0; i < 100; i++ {
		ns[i] = N{
			foo: randomSliceOfString(5),
			bar: randomSliceOfInt64(5),
			a: a,
		}

	}

	b.ResetTimer()
	for test := 0; test < b.N; test++ {
		for _, n := range ns {
			n.Calc()	
		}
	}
}

func BenchmarkParameterInStructAsPointer(b *testing.B) {
	a := &A{
		x: "a",
		y: 1,
	}

	ns := make([]NPointer, 100)
	for i := 0; i < 100; i++ {
		ns[i] = NPointer{
			foo: randomSliceOfString(5),
			bar: randomSliceOfInt64(5),
			a: a,
		}

	}

	b.ResetTimer()
	for test := 0; test < b.N; test++ {
		for _, n := range ns {
			n.Calc()	
		}
	}
}

// func BenchmarkFuncParamsPointer(b *testing.B) {
// 	i := A{
// 		foo: randomSliceOfString(5),
// 		bar: randomSliceOfInt64(5),
// 	}

// 	j := A{
// 		foo: randomSliceOfString(5),
// 		bar: randomSliceOfInt64(5),
// 	}

// 	b.ResetTimer()
// 	for test := 0; test < b.N; test++ {
// 		mergeStructsPointer(&i, &j)
// 	}
// }

const charset = "abcdefghijklmnopqrstuvwxyz"

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomSliceOfString(resultLen int) []string {
	s := make([]string, resultLen)
	for i := range s {
		b := make([]byte, 5)
		for j := range b {
			b[j] = charset[seed.Intn(len(charset))]
		}
		s[i] = string(b)
	}

	return s
}

func randomSliceOfInt64(len int) []int64 {
	s := make([]int64, len)
	for i := range s {
		s[i] = seed.Int63()
	}

	return s
}
