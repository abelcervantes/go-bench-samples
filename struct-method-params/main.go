package main

import (
	"math/rand"
	"time"

	"github.com/pkg/profile"
)

type Modifier struct {
	modifier1 []int64
	modifier2 []int64
	modifier3 bool
	modifier4 int64
}

type N struct {
	foo []int64
	bar []int64
	mod Modifier
	m   int64
}

func (n *N) Calc() {
	var m int64 = 0
	for i := range n.foo {
		n.foo[i] = n.foo[i] - n.mod.modifier1[i]
		m += n.foo[i]
	}
	for i := range n.bar {
		n.bar[i] = n.bar[i] - n.mod.modifier2[i]
		m += n.bar[i]
	}

	if n.mod.modifier3 {
		m -= n.mod.modifier4
	}

	n.m = m
}

type NPointer struct {
	foo []int64
	bar []int64
	mod *Modifier
	m   int64
}

func (n *NPointer) Calc() {
	var m int64 = 0
	for i := range n.foo {
		n.foo[i] = n.foo[i] - n.mod.modifier1[i]
		m += n.foo[i]
	}
	for i := range n.bar {
		n.bar[i] = n.bar[i] - n.mod.modifier2[i]
		m += n.bar[i]
	}

	if n.mod.modifier3 {
		m -= n.mod.modifier4
	}

	n.m = m
}

type NFunc struct {
	foo []int64
	bar []int64
	m   int64
}

func (n *NFunc) Calc(mod Modifier) {
	var m int64 = 0
	for i := range n.foo {
		n.foo[i] = n.foo[i] - mod.modifier1[i]
		m += n.foo[i]
	}
	for i := range n.bar {
		n.bar[i] = n.bar[i] - mod.modifier2[i]
		m += n.bar[i]
	}

	if mod.modifier3 {
		m -= mod.modifier4
	}

	n.m = m
}

func (n *NFunc) CalcPointer(mod *Modifier) {
	var m int64 = 0
	for i := range n.foo {
		n.foo[i] = n.foo[i] - mod.modifier1[i]
		m += n.foo[i]
	}
	for i := range n.bar {
		n.bar[i] = n.bar[i] - mod.modifier2[i]
		m += n.bar[i]
	}

	if mod.modifier3 {
		m -= mod.modifier4
	}

	n.m = m
}

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomSliceOfInt64(len int) []int64 {
	s := make([]int64, len)
	for i := range s {
		s[i] = seed.Int63()
	}

	return s
}

var l = 2000000

func methodParamAsValue() {
	mod := Modifier{
		modifier1: randomSliceOfInt64(20),
		modifier2: randomSliceOfInt64(20),
		modifier3: true,
		modifier4: 10,
	}

	ns1 := make([]NFunc, l)
	for i := 0; i < l; i++ {
		ns1[i] = NFunc{
			foo: randomSliceOfInt64(20),
			bar: randomSliceOfInt64(20),
		}
	}
	for _, n := range ns1 {
		n.Calc(mod)
	}
}

func methodParamAsPointer() {
	mod := &Modifier{
		modifier1: randomSliceOfInt64(20),
		modifier2: randomSliceOfInt64(20),
		modifier3: true,
		modifier4: 10,
	}

	ns1 := make([]NFunc, l)
	for i := 0; i < l; i++ {
		ns1[i] = NFunc{
			foo: randomSliceOfInt64(20),
			bar: randomSliceOfInt64(20),
		}
	}
	for _, n := range ns1 {
		n.CalcPointer(mod)
	}
}

func paramInStructAsValue() {
	mod := Modifier{
		modifier1: randomSliceOfInt64(20),
		modifier2: randomSliceOfInt64(20),
		modifier3: true,
		modifier4: 10,
	}

	ns1 := make([]N, l)
	for i := 0; i < l; i++ {
		ns1[i] = N{
			foo: randomSliceOfInt64(20),
			bar: randomSliceOfInt64(20),
			mod: mod,
		}
	}
	for _, n := range ns1 {
		n.Calc()
	}
}

func paramInStructAsPointer() {
	mod := &Modifier{
		modifier1: randomSliceOfInt64(20),
		modifier2: randomSliceOfInt64(20),
		modifier3: true,
		modifier4: 10,
	}

	ns1 := make([]NPointer, l)
	for i := 0; i < l; i++ {
		ns1[i] = NPointer{
			foo: randomSliceOfInt64(20),
			bar: randomSliceOfInt64(20),
			mod: mod,
		}
	}
	for _, n := range ns1 {
		n.Calc()
	}
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("./pprof/param_in_struct_pointer")).Stop()
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath("./pprof/method_param_pointer")).Stop()
	//methodParamAsValue()
	//methodParamAsPointer()
	//paramInStructAsValue()
	paramInStructAsPointer()
}
