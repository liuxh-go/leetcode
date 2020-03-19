package simple

/*
	给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

	在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

	注意:
	假设字符串的长度不会超过 1010。

	思路:
	只需要统计字符串中出现次数为偶数次的字符个数,如果有奇数次的出现,加最多奇数次的
*/

func longestPalindrome(s string) int {
	charMap := make([]int, 52)

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			charMap[c-'a']++
		}

		if c >= 'A' && c <= 'Z' {
			charMap[c-'A'+26]++
		}
	}

	result := 0
	for _, count := range charMap {
		if count%2 == 0 {
			result += count
		}
		if count%2 == 1 {
			result += count - 1
		}
	}

	if result < len(s) {
		result++
	}

	return result
}
