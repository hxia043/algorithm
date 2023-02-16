package main

import "fmt"

const n int = 6

func selectionSort(data [n]int) {
	for i := 0; i < len(data)-1; i++ {
		min := data[i]
		j := i + 1
		selected := j

		for ; j < len(data); j++ {
			if min > data[j] {
				min = data[j]
				selected = j
			}
		}

		data[selected] = data[i]
		data[i] = min
	}

	fmt.Println(data)
}

func main() {
	var A = [6]int{4, 5, 6, 3, 1, 2}
	selectionSort(A)
}
