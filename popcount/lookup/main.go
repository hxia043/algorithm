package main

import (
	"fmt"
)

func main() {
	/*
		// v1.0
		table := make(map[uint32]int)
		table[0] = 0
		table[1] = 1
		table[2] = 1
		table[3] = 2
		table[4] = 1
		table[5] = 2
		table[6] = 2
		table[7] = 3
		table[8] = 1
		table[9] = 2
		table[10] = 2
		table[11] = 3
		table[12] = 2
		table[13] = 3
		table[14] = 3
		table[15] = 4

		var n uint32 = 100
		n1 := n & 0xf000 >> 12
		n2 := n & 0x0f00 >> 8
		n3 := n & 0x00f0 >> 4
		n4 := n & 0x000f

		fmt.Printf("%08b\n", n)
		fmt.Printf("%08b\n", n1)
		fmt.Printf("%08b\n", n2)
		fmt.Printf("%08b\n", n3)
		fmt.Printf("%08b\n", n4)

		fmt.Println(table[n1] + table[n2] + table[n3] + table[n4])
	*/

	// v1.1
	var n uint32 = 100

	lookup := []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}
	count := 0

	for n != 0 {
		count += lookup[n&0xF]
		n = n >> 4
	}

	fmt.Println(count)
}
