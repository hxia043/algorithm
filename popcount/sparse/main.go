package main

import "fmt"

var n uint64 = 16

func main() {
	count := 0

	for n != 0 {
		count++
		n &= n - 1
	}

	fmt.Println(count)
}
