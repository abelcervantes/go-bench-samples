package main

import (
	"testing"
)

var randomSliceSize = 20
var randomSlice1 = randomSliceOfInt64(randomSliceSize)
var randomSlice2 = randomSliceOfInt64(randomSliceSize)
var randomSlice3 = randomSliceOfInt64(randomSliceSize)
var randomSlice4 = randomSliceOfInt64(randomSliceSize)

func BenchmarkMethodParamAsValue(b *testing.B) {
	mod := Modifier{
		modifier1: randomSlice3,
		modifier2: randomSlice4,
		modifier3: true,
		modifier4: 10,
	}

	n := NFunc{
		foo: randomSlice1,
		bar: randomSlice2,
	}

	b.ResetTimer()
	for test := 0; test < b.N; test++ {
		n.Calc(mod)
	}
}

func BenchmarkMethodParamsAsPointer(b *testing.B) {
	mod := &Modifier{
		modifier1: randomSlice3,
		modifier2: randomSlice4,
		modifier3: true,
		modifier4: 10,
	}

	n := NFunc{
		foo: randomSlice1,
		bar: randomSlice2,
	}

	b.ResetTimer()
	for test := 0; test < b.N; test++ {
		n.CalcPointer(mod)
	}
}

func BenchmarkParameterInStructAsValue(b *testing.B) {
	mod := Modifier{
		modifier1: randomSlice3,
		modifier2: randomSlice4,
		modifier3: true,
		modifier4: 10,
	}

	n := N{
		foo: randomSlice1,
		bar: randomSlice2,
		mod: mod,
	}

	b.ResetTimer()
	for test := 0; test < b.N; test++ {
		n.Calc()
	}
}

func BenchmarkParameterInStructAsPointer(b *testing.B) {
	mod := &Modifier{
		modifier1: randomSlice3,
		modifier2: randomSlice4,
		modifier3: true,
		modifier4: 10,
	}

	n := NPointer{
		foo: randomSlice1,
		bar: randomSlice2,
		mod: mod,
	}

	b.ResetTimer()
	for test := 0; test < b.N; test++ {
		n.Calc()
	}
}
