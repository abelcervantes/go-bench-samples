package samples

import (
	"fmt"
	"log"
	"strings"
)

type A struct {
	x string
	y int64
}

type N struct {
	foo []string
	bar []int64
	a A
}

func (n *N) Calc() {
	for i := range n.foo {
		n.foo[i] = strings.Replace(n.foo[i], n.a.x, "X", -1)
	}
	for i := range n.bar {
		n.bar[i] = n.bar[i] - n.a.y
	}
}

type NPointer struct {
	foo []string
	bar []int64
	a *A
}

func (n *NPointer) Calc() {
	for i := range n.foo {
		n.foo[i] = strings.Replace(n.foo[i], n.a.x, "X", -1)
	}
	for i := range n.bar {
		n.bar[i] = n.bar[i] - n.a.y
	}
}

// func mergeStructs(i, j A) A {
// 	return A{
// 		foo: mergeSlicesOfStrings(i.foo, j.foo),
// 		bar: mergeSlicesOfInt64(i.bar, j.bar),
// 	}
// }

func mergeSlicesOfStrings(i, j []string) []string {
	if len(i) != len(j) {
		log.Fatal("slices of string dif size")
	}

	result := make([]string, len(i))
	for x := 0; x < len(i); x++ {
		result[x] = fmt.Sprintf("%s.%s", i[x], j[x])
	}

	return result
}

func mergeSlicesOfInt64(i, j []int64) []int64 {
	if len(i) != len(j) {
		log.Fatal("slices of int64 dif size")
	}

	result := make([]int64, len(i))
	for x := 0; x < len(i); x++ {
		result[x] = i[x] - j[x]
	}

	return result
}

// func mergeStructsPointer(i, j *A) A {
// 	return A{
// 		foo: mergeSlicesOfStrings(i.foo, j.foo),
// 		bar: mergeSlicesOfInt64(i.bar, j.bar),
// 	}
// }

func mergeSlicesOfStringsPointer(i, j *[]string) []string {
	if len(*i) != len(*j) {
		log.Fatal("slices of string dif size")
	}

	result := make([]string, len(*i))
	for x := 0; x < len(*i); x++ {
		result[x] = fmt.Sprintf("%s.%s", (*i)[x], (*j)[x])
	}

	return result
}

func mergeSlicesOfInt64Pointer(i, j *[]int64) []int64 {
	if len(*i) != len(*j) {
		log.Fatal("slices of int64 dif size")
	}

	result := make([]int64, len(*i))
	for x := 0; x < len(*i); x++ {
		result[x] = (*i)[x] - (*j)[x]
	}

	return result
}