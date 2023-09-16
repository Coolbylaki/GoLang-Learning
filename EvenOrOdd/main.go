package main

import "fmt"

func main() {
	numbers := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("Number %v is even\n", num)
		} else {
			fmt.Printf("Number %v is odd\n", num)
		}
	}
}
