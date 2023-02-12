package main

import "fmt"

const n int = 6

var isBubbled bool

func bubbleWithoutFlag(data [n]int) [n]int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}

		fmt.Println(data)
	}

	return data
}

func bubbleWithFlag(data [n]int) [n]int {
	for i := 0; i < len(data)-1; i++ {

		isBubbled = false

		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isBubbled = true
			}
		}

		fmt.Println(data)

		if !isBubbled {
			break
		}
	}

	return data
}

func main() {
	var A = [6]int{3, 5, 4, 1, 2, 6}

	bubbleWithoutFlag(A)

	fmt.Println("=====================")

	bubbleWithFlag(A)
}
