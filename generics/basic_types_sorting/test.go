package main

import (
    "fmt"
    "sort"
)

type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}

type sortable[T Ordered] []T

func (s sortable[T]) Len() int {
    return len(s)
}

func (s sortable[T]) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortable[T]) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func Sort[T Ordered](a []T) {
    sort.Sort(sortable[T](a))
}

func main() {
    ints := []int{1, 3, 2, 5, 4}
    Sort(ints)
    fmt.Println(ints) // [1 2 3 4 5]

    floats := []float64{1.30, 3.20, 2.10, 5.40, 4.50}
    Sort(floats)
    fmt.Println(floats) // [1.3 2.1 3.2 4.5 5.4]

    strings := []string{"a", "e", "b", "d", "c"}
    Sort(strings)
    fmt.Println(strings) // [a b c d e]
}
