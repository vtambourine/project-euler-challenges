package main

import "fmt"

func Ints(n int) []int {
	result := []int{}
	var d, m int
	d = n

	if (d == 0) {
		return []int{0}
	}

	for d != 0 {
		m = d % 10
		d = d / 10
		result = append([]int{m}, result...)
	}
	return result
}

func main() {
	a := Ints(100)
	fmt.Println(a)
	fmt.Println(len(a))
}