package sparse

//var n uint64 = 16

func Sparse(n uint64, count uint64) {
	for n != 0 {
		count++
		n &= n - 1
	}
}

/*
func main() {
	var count uint64
	sparse(15, count)
}
*/
