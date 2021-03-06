package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// N  Only primes less than or equal to N will be generated
const N = 1001

func main() {
	var err error
	var x, y, n int

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <Max_Prime>\n", os.Args[0])
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}

	nsqrt := math.Sqrt(float64(N))

	isPrime := make([]bool, N)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < N; y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
			fmt.Println(x)
		}
	}

	// primes is now a slice that contains all the
	// primes numbers up to N

	// let's print them
	// for _, x := range primes {
	// 	fmt.Println(x)
	// }
}
