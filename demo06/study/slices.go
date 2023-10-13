package study

import (
	"cmp"
	"fmt"
	"slices"
)

func TestSort() {
	strs := []string{"c", "a", "b"}

	slices.Sort(strs)
	fmt.Println("strs:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("ints:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("sorted:", s)
}

func TestCustomSort() {
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits:", fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("people:", people)
}
