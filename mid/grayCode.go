package mid

/*
	格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。

	给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。

	ex:
	输入: 2
	输出: [0,1,3,2]
	解释:
	00 - 0
	01 - 1
	11 - 3
	10 - 2

	输入:3
	输出:[0,]
	解释:
	000 - 0
	001 - 1
	011	- 3
	010 - 2
	110 - 6
	100 - 4
	101 - 5
	111 - 7

	背答案题:
	关键是搞清楚格雷编码的生成过程, G(i) = i ^ (i/2);
	如 n = 3:
	G(0) = 000,
	G(1) = 1 ^ 0 = 001 ^ 000 = 001
	G(2) = 2 ^ 1 = 010 ^ 001 = 011
	G(3) = 3 ^ 1 = 011 ^ 001 = 010
	G(4) = 4 ^ 2 = 100 ^ 010 = 110
	G(5) = 5 ^ 2 = 101 ^ 010 = 111
	G(6) = 6 ^ 3 = 110 ^ 011 = 101
	G(7) = 7 ^ 3 = 111 ^ 011 = 100
*/

func grayCode(n int) []int {
	l := 1 << n
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = i ^ i>>1
	}

	return result
}
