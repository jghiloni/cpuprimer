package main

import (
	"fmt"
	"time"
)

const max = ^uint64(0)

func main() {
	time.Sleep(10 * time.Second)
	for i := uint64(2); i <= max; i++ {
		fmt.Printf("is %d prime? %t", i, isPrime(i))
	}
}

func isPrime(num uint64) bool {
	var divisor uint64
	for divisor = 2; divisor < num; divisor++ {
		if num == divisor {
			continue
		}

		if num%divisor == 0 {
			return false
		}
	}

	return true
}
