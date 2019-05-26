package main

import (
	"fmt"
	"time"
)

func sieve(num int) {

	// do not use
	// prime := [num+1]bool
	// will cause : non-constant array bound num + 1 error

	// an array of boolean - the idiomatic way
	prime := make([]bool, num+1)

	// initialize everything with false first(not crossed)
	for i := 0; i < num+1; i++ {
		prime[i] = false
	}

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true // cross
			}
		}
	}
}

func main() {
	start := time.Now()
	sieve(100000)
	fmt.Println("Took", time.Since(start))
}
