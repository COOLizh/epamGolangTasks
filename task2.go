package main

import (
	"fmt"
	"strconv"
)

func factorial(i uint) uint {
	var result uint = 1
	for j := uint(2); j <= i; j++ {
		result *= j
	}
	return result
}

func main() {
	var possibleInput string
	for {
		fmt.Print("Enter n: ")
		fmt.Scan(&possibleInput)

		u64, err := strconv.ParseUint(possibleInput, 10, 32)
		if err != nil {
			fmt.Println("Check your input. n must be uint type. Try again...")
			continue
		}

		n := uint(u64)
		if n > 20 {
			fmt.Println("Please enter a variable from 0 to 20, otherwise overflow will occur in the factorial function. Try again...")
			continue
		}

		fmt.Println("Factorial ", n, " is ", factorial(n))
		break
	}
}
