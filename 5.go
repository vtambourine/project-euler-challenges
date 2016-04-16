package main

import (
	"fmt"
	"math"
)

// const N = float64(2520)

func main() {
	// var i float64
	// for i = 1; i <= 10; i++ {
	// 	fmt.Printf("%v; %v\n", N / i, math.Mod(N, i))
	// }

	fmt.Println("Starting...")

	const N = 20
	var (
		i float64 = 1
		n int = N
		done bool = false
	)
	for {
		n++
		for i = N; i > 1; i-- {
			// fmt.Printf("Trying %v / %v... ", n, i)
			if (math.Mod(float64(n), i) != 0) {
				// fmt.Printf("Not %v\n", n)
				break
			}
			
			if (i == 2) {
				// fmt.Printf("\nAnd we are done with %v\n", n)
				done = true	
			}
		}
		if (done) {
			break
		}
	}
	fmt.Println()
	fmt.Println(n)
}