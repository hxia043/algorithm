# popcount
popcount 查找二进制中 1 的数目。  

对于十进制整数可通过除 2 取余法获得该整数的二进制数。那么，对于二进制中 0 和 1 的操作，根据不同场景会有不同算法。popcount 算法用来提取二进制中 1 的数目。具体使用场景有待探索。

## interated popcount
很自然的想到用遍历来实现，设置计数器，遍历所有二进制位数，累加，得到计数器的值。示例如下：
```
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
```

这是 1.0 版本的遍历，可以看出还有改进的点：
- 循环条件可放在 for 循环体中。
- `(n & 1)` 得到的结果不是 0 就是 1。且，为 1 计数器加 1。可以根据这一隐含条件将 count 和 `(n & 1)` 统一成一条语句。优点是更简洁，缺点是每次比一下都要加 count，即使比的结果为 `0x0`。

1.1 版本遍历如下：
```
var count uint64

for ; n != 0; n = n >> 1 {
    count += (n & 1)
}

fmt.Println(count)
```

## sparse popcount
遍历的时间复杂度为 O(bit)，当 bit 位数很大时该算法就不太适用了。

通过科学归纳法得出通过 `n &= n - 1` 可以消除 n 末位的 1。即，消除了几个 1 就得到了总的二进制位 1 的个数。图解如下：  

![sparse](sparse/img/sparse.png "sparse")

示例代码如下：
```
count := 0

for n != 0 {
    count++
    n &= n - 1
}

fmt.Println(count)
```

这个算法已经相当棒了，时间复杂度从 O(bit)，降到了 O(1-bit)，1-bit 表示二进制位中 bit 位 1 的个数。对于只有 1 个二进制位的 n，时间复杂度仅为 O(1)。

既然有 sparse 就应该有 dense，是的。dense 是反着来，当二进制位有多个 1 时，对该二进制位取反再 sparse 一下就能减少不少运算，得到的结果被二进制位一减即得到二进制的个数。具体示例不列了，可参考 [popcount 算法分析](https://zhuanlan.zhihu.com/p/341488123)。

## lookup popcount
`lookup` 的设想是，可以分批次统计二进制位中 1 的个数，将每一批次值相加即可得总的二进制 1 的位数。

即是分批次，那可以通过循环迭代和查表的方式来实现查询。循环迭代效果不高，而查表则是空间换时间。这里用查表来实现 `lookup popcount`。示例代码如下：
```
var n uint32 = 100

lookup := []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}
count := 0

for n != 0 {
    count += lookup[n&0xF]
    n = n >> 4
}

fmt.Println(count)
```

相比于 1.0 版查表，这里最重要的改变是，通过切片而不是 map 构造表。切片隐含的一层意思是值对应二进制位个数。如整数 15 通过 lookup[15] 即可从切片中查到其二进制位 1 个数是 4。

## 性能测试
测试三种 popcount 算法，性能如下:
```
# go test -bench=. -v -benchmem -benchtime=5s -count=3 .
goos: linux
goarch: amd64
pkg: popcount/test
cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
BenchmarkIterated
BenchmarkIterated-3     195811508               36.39 ns/op            0 B/op          0 allocs/op
BenchmarkIterated-3     180319236               31.16 ns/op            0 B/op          0 allocs/op
BenchmarkIterated-3     171637110               33.30 ns/op            0 B/op          0 allocs/op
BenchmarkSparse
BenchmarkSparse-3       271943863               22.12 ns/op            0 B/op          0 allocs/op
BenchmarkSparse-3       268430272               23.56 ns/op            0 B/op          0 allocs/op
BenchmarkSparse-3       244256011               22.03 ns/op            0 B/op          0 allocs/op
BenchmarkLookup
BenchmarkLookup-3       504271546               11.19 ns/op            0 B/op          0 allocs/op
BenchmarkLookup-3       554806210               10.33 ns/op            0 B/op          0 allocs/op
BenchmarkLookup-3       711143409               12.16 ns/op            0 B/op          0 allocs/op
PASS
ok      popcount/test   76.752s
```

## 结尾
当然 popcount 并不只有这几种，且 Go 中常用的是 [hacker popcount](https://zhuanlan.zhihu.com/p/341488123)。go 中锁的代码包含大量的二进制位操作，理解 popcount 也能对后续理解锁有一点帮助。  

尝试去理解 `hacker popcount` 并没有理解透，这里先记录以上几种。后续有待更新。

## 引用
- [hacker popcount](https://zhuanlan.zhihu.com/p/341488123)
- [go popcount](https://github.com/tmthrgd/go-popcount)
