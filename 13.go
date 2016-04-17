package main

import (
	"bufio"
	"fmt"
	"os" // "math"
	"strconv"
)

func Ints(n int) []int {
	result := []int{}

	var d, m int
	d = n

	if d == 0 {
		return []int{0}
	}

	for d != 0 {
		m = d % 10
		d = d / 10
		result = append([]int{m}, result...)
	}
	return result
}

func Sum(array []int) []int {
	var l int = 0
	for i := len(array) - 1; i >= 0; i-- {
		v := array[i] + l
		array[i] = v % 10
		l = v / 10
	}
	// fmt.Println(Ints(l))
	// fmt.Println(array)
	result := append(Ints(l), array...)
	return result
}

func main() {
	fmt.Println("Expect: ", 99999*16)
	file, _ := os.Open("13.input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	// var s string
	a := [][]int{}
	var r int = 0
	row := []int{}
	var n int64
	for scanner.Scan() {
		s := scanner.Text()
		if s == "\n" {
			a = append(a, row)
			r++
			row = []int{}
		} else {
			n, _ = strconv.ParseInt(s, 10, 0)
			row = append(row, int(n))
		}
	}

	result := []int{}
	cn := 50
	rn := 100
	var sum int = 0

	// iterate over columns in backward order
	for i := cn - 1; i >= 0; i-- {
		// sum of current position
		sum = 0
		for j := 0; j < rn; j++ {
			// fmt.Println(j, i, a[0][4])
			// a_ji — current digit
			sum = sum + a[j][i]
		}
		// fmt.Printf("%v\n", []byte(strconv.FormatInt(int64(sum), 10)))
		// fmt.Println(sum)

		// sd := Ints(sum)
		// fmt.Println(i, sd)

		result = append([]int{sum}, result...)
	}
	// result[0] = 5
	// fmt.Println(result)
	ss := Sum(result)
	fmt.Println(ss)
	fmt.Println()
	fmt.Println(ss[:10])
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", ss[i])
	}

	// for _, x := range a {
	// 	fmt.Printf("%v: %T\n", x, x)
	// 	for i := len(x)
	// }
}
