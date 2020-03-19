package mid

/*
	一条包含字母 A-Z 的消息通过以下方式进行了编码：

	'A' -> 1
	'B' -> 2
	...
	'Z' -> 26
	给定一个只包含数字的非空字符串，请计算解码方法的总数。

	ex:
	输入: "226"
	输出: 3
	解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
*/

func numDecodings(s string) int {
	l := len(s)
	if l == 0 || (l == 1 && s[0] == '0') {
		return 0
	}

	a, b := 1, 1
	for i := 0; i < l; i++ {
		temp := 0
		if s[i] != '0' {
			temp += b
		}

		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6')) {
			temp += a
		}

		a, b = b, temp
	}

	return b
}
