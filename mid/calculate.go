package mid

import (
	"strconv"

	"wshhz.com/leetcode/stack"
)

/*
	实现一个基本的计算器来计算一个简单的字符串表达式的值。

	字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

	利用 栈 + 递归 的思路解决
*/

func calculate(s string) int {
	_, result := dfs(s, 0)

	return result
}

func dfs(s string, index int) (int, int) {
	stackObj := stack.NewStack()

	num := 0
	sign := uint8('+')
	for ; index < len(s); index++ {
		c := s[index]

		// 计数
		if c >= '0' && c <= '9' {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			num = num*10 + digit
		}

		// 递归计算括号中的内容
		if c == '(' {
			index, num = dfs(s, index+1)
		}

		if ((c < '0' || c > '9') && c != ' ') || index == len(s)-1 {
			switch sign {
			case '+':
				stackObj.Push(num)
			case '-':
				stackObj.Push(-num)
			case '*':
				// 栈顶元素弹出,计算和接下来的数的乘积
				data := stackObj.Pop()
				if data == nil {
					panic("stack nil!")
				}
				stackObj.Push(*data * num)
			case '/':
				// 栈顶元素弹出,计算和接下来的数的商
				data := stackObj.Pop()
				if data == nil {
					panic("stack nil!")
				}
				stackObj.Push(*data / num)
			}

			sign = c
			num = 0
		}

		// 递归结束
		if c == ')' {
			break
		}
	}

	return index, stackObj.Sum()
}
