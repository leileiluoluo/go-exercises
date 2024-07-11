package main

import (
    "fmt"
    "sort"
)

type Comparable[T any] interface {
    CompareTo(T) int
}

type sortable[T Comparable[T]] []T

func (s sortable[T]) Len() int {
    return len(s)
}

func (s sortable[T]) Less(i, j int) bool {
    return s[i].CompareTo(s[j]) < 0
}

func (s sortable[T]) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

type student struct {
    id   int
    name string
}

func (s student) CompareTo(other student) int {
    return s.id - other.id
}

func Sort[T Comparable[T]](a []T) {
    sort.Sort(sortable[T](a))
}

func main() {
    students := []student{
        {id: 1, name: "Larry"},
        {id: 3, name: "Jacky"},
        {id: 2, name: "Lucy"},
    }

    Sort(students)

    fmt.Println(students) // [{1 Larry} {2 Lucy} {3 Jacky}]
}
