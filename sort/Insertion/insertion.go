package main

import (
	"fmt"
	"os"
)

const n int = 6

func insertPositiveSort(data [n]int) error {
	if len(data) == 0 {
		return fmt.Errorf("empty data")
	}

	if len(data) == 1 {
		fmt.Println(data)
	}

	for j := 1; j < len(data); j++ {
		for i := 0; i < j; i++ {
			if data[i] > data[j] {
				temp := data[j]
				for g := 0; g < j-i; g++ {
					data[j-g] = data[j-g-1]
				}
				data[i] = temp

				break
			}
		}
	}

	fmt.Println(data)

	return nil
}

func insertReverseSort(data [6]int) error {
	if len(data) == 0 {
		return fmt.Errorf("empty data")
	}

	if len(data) == 1 {
		fmt.Println(data)
	}

	for j := 1; j < len(data); j++ {
		value := data[j]
		i := j - 1

		for ; i >= 0; i-- {
			if data[i] < value {
				break
			}

			data[i+1] = data[i]
		}

		data[i+1] = value
	}

	fmt.Println(data)
	return nil
}

func main() {
	var A = [6]int{4, 5, 6, 3, 1, 2}

	if err := insertPositiveSort(A); err != nil {
		os.Exit(1)
	}

	var B = [6]int{4, 5, 6, 3, 1, 2}
	if err := insertReverseSort(B); err != nil {
		os.Exit(1)
	}
}
