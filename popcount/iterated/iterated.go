package iterated

//var n uint64 = 15

func Iterated(n uint64, count uint64) {
	for ; n != 0; n = n >> 1 {
		count += (n & 1)
	}
}

/*
func main() {
	// v1.0

		count := 0

		for {
			if (n & 1) == 1 {
				count++
			}
			n = n >> 1

			if n == 0 {
				break
			}
		}

		fmt.Println(count)


	// v1.1
	var count uint64
	iterated(15, count)
}
*/
