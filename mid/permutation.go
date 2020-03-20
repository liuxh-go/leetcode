package mid

import "sort"

/*
	有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。

	示例1:
	 输入：S = "qqe"
	 输出：["eqq","qeq","qqe"]
*/

func permutation(S string) []string {
	length := len(S)
	used := make([]bool, length)
	result := make([]string, 0)

	data := []byte(S)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	var dfsFunc func(s string)
	dfsFunc = func(s string) {
		if len(s) == length {
			result = append(result, s)
			return
		}

		for i := 0; i < length; i++ {
			if used[i] == true || (i > 0 && data[i] == data[i-1] && used[i-1] == false) {
				continue
			}

			used[i] = true
			dfsFunc(s + string(data[i]))
			used[i] = false
		}
	}

	dfsFunc("")

	return result
}
