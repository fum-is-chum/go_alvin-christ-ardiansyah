package main

import (
	"fmt"
)

var primes []int

func isPrimeNumber(number int) bool {
	if number < 2 {
		return false
	} else {
		for i:= 2; i * i <= number; i++ {
			if number % i == 0 {
				return false;
			}
		}
		return true;
	}
}

func primeX(number int) int {
	lastPrimeIndex := len(primes)

	if len(primes) >= number {
		return primes[number - 1]
	}

	primeToFind := number - lastPrimeIndex // how many primes left to find
	currentPrime := primes[lastPrimeIndex - 1] // set currentPrime to last found prime

	for primeToFind > 0 {								
		currentPrime++;
		if(isPrimeNumber(currentPrime)) {
			primes = append(primes, currentPrime)
			primeToFind--
		}
	}

	return primes[number-1]
}

func main() {

	// store prime
	primes = []int{2,3,5,7,11}

	fmt.Println(primeX(101)) // 547
	
	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29


}
