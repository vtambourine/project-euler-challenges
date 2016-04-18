package main

import (
	"fmt"
)

func main() {
	limit := int(2e6)
	sieve := [2e6]bool{true, true} // true if composite
	var i int
	for n := 2; n < limit; n++ {
		if sieve[n] {
			continue
		}
		// fmt.Println("is prime", n)
		i = n * n
		for i < limit {
			// fmt.Println("for", n, "mark", i)
			sieve[i] = true
			i = n + i
		}
	}

	sum := 0
	for k, v := range sieve {
		// fmt.Println(k, v)		
		if !v {
			sum += k
		}
	}

	fmt.Println("sieve", sum)
}