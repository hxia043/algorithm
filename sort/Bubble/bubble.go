package main

import "fmt"

func main() {
	var A = [6]int{10, 4, 5, 0, 2, 1}

	for i := 0; i < len(A)-1; i++ {
		for j := 0; j < len(A)-1-i; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}

	fmt.Println(A)
}
