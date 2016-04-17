package main

import (
	"fmt"
	"math/big"
	"math"
)

var TARGET int64 = 600851475143
// var TARGET = big.NewInt(600851475143)
// var TARGET = big.NewInt(22)

func main() {
	fmt.Println("hello")

	n := int64(math.Sqrt(float64(TARGET)))

	for n > 0 {
		if (TARGET % n == 0) && big.NewInt(n).ProbablyPrime(20) {
			break
		}
		n--
	}
	
	fmt.Println(n)

	// var a int64 = 1
	// // inc := big.NewInt(1)
	// var limit int64 
	// limit = TARGET / 2
	// // limit := big.NewInt(21)
	// // zero := big.NewInt(0)
	// var mod int64
	// var max int64
	// for a < limit {
	// 	a++
	// 	if (big.NewInt(int64(a)).ProbablyPrime(20)) {
	// 		// fmt.Println("Is prime", a)
	// 		mod = TARGET % a
	// 		// fmt.Println(TARGET, "mod", a, "=", mod)
	// 		// fmt.Println("?? =", mod.Cmp(zero) == 0)
	// 		if mod == 0 {
	// 			max = a
	// 			fmt.Println(">>>> Set to", a)
	// 		}
	// 	}
	// }
	// fmt.Println("MAX =", max)
	// fmt.Println("?? =", max.Cmp(zero) == 0)
	// fmt.Println(a.Add(a, big.NewInt(30)))
	// fmt.Println(big.DivMod(TARGET, limit))
}
