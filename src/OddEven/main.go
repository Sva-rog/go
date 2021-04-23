package main

import "fmt"

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, nums := range num {
		if num[i]%2 == 0 {
			fmt.Printf("Even! Position: %v Number: %v\n", i, nums)
		} else {
			fmt.Printf("Odd! Position: %v Number: %v\n", i, nums)
		}

	}
}
