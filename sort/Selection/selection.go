package main

import "fmt"

func main() {
	var A = [6]int{4, 5, 6, 3, 1, 2}

	for i := 0; i < len(A)-1; i++ {
		min := A[i]
		j := i + 1
		selected := j

		for ; j < len(A); j++ {
			if min > A[j] {
				min = A[j]
				selected = j
			}
		}

		A[selected] = A[i]
		A[i] = min
	}

	fmt.Println(A)
}
