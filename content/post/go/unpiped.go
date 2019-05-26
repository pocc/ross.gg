package main

import "fmt"
import "time"

func sieve(primes []int, index *int) int {
	for {
		isPrime := true
		for _, prime := range primes {
			if *index%prime == 0 {
				isPrime = false
				break
			} 
		}
		if isPrime {
			return *index
		}
		*index++
	}
}

func main() {
	start := time.Now()
	numPrimes := 10000
	primes := make([]int, numPrimes)
	// For each number, divide by previous primes
	index := 2 
	for i:=0; i<numPrimes; i++ {
		primes[i] = sieve(primes[:i], &index)
		//fmt.Println(primes[i])
	}
	fmt.Println("Took", time.Since(start))
}
