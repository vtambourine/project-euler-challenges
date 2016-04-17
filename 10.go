package main

import (
	"fmt"
	"math/big"
)



func main() {
	var b *big.Int
	sum := 2
	limit := int(2e6)
	for i := 3; i <= limit; i = i + 2 {
		b = big.NewInt(int64(i))
		if b.ProbablyPrime(20) {
			sum += i
		}
	}
	fmt.Print(sum)
}
