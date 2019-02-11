package main

import "fmt"



func main() {
	numSlice := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range numSlice {
		if numSlice[i]%2 == 0 {
			fmt.Printf("%v is even\n", numSlice[i])
		} else {
			fmt.Printf("%v is odd", numSlice[i])
		}
	}
}
