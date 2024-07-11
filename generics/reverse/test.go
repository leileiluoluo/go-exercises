package main

import (
    "fmt"
)

func ReverseInts(a []int) {
    for i, j := 0, len(a)-1; i < j; {
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }
}

func ReverseFloat64s(a []float64) {
    for i, j := 0, len(a)-1; i < j; {
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }
}

func ReverseStrings(a []string) {
    for i, j := 0, len(a)-1; i < j; {
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }
}

type student struct {
    id   int
    name string
}

func ReverseStudents(a []student) {
    for i, j := 0, len(a)-1; i < j; {
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }
}

func Reverse[T any](a []T) {
    for i, j := 0, len(a)-1; i < j; {
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }
}

func main() {
    ints := []int{1, 2, 3, 4, 5}
    Reverse(ints)
    fmt.Println(ints) // [5 4 3 2 1]

    floats := []float64{1.03, 2.25, 3.38, 4.49, 5.52}
    Reverse(floats)
    fmt.Println(floats) // [5.52 4.49 3.38 2.25 1.03]

    strings := []string{"a", "b", "c", "d", "e"}
    Reverse(strings)
    fmt.Println(strings) // [e d c b a]

    students := []student{
        {id: 1, name: "Larry"},
        {id: 2, name: "Jacky"},
        {id: 3, name: "Alice"},
        {id: 4, name: "Lucy"},
        {id: 5, name: "Cindy"},
    }
    Reverse(students)
    fmt.Println(students) // [{5 Cindy} {4 Lucy} {3 Alice} {2 Jacky} {1 Larry}]
}
