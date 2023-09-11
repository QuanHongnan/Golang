/*
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
 

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
 

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

package main
import "fmt"

func isValid(s string)bool{
	n := len(s)
	if n%2==1{
		return false
	}

	pairs := map[byte]byte{//右括号 为键 左括号为 值
		'}':'{',
		')':'(',
		']':'[',
	}

	stack := []byte{}
	for i:=0; i < n ;i++{
		if pairs[s[i]]>0{ //看看是不是右括号 如果是向下运行 不是则压入栈中
			//本次循环到右括号 在栈中找相对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]]{ //如果栈中没有值了或者栈中对应的值并不是
				return false
			}
			//栈中有值并且栈中的值对应本次循环的左括号 则将此括号剔除栈中 开始下一轮循环
			fmt.Printf("%s \n",stack)
			stack = stack[:len(stack)-1]

		}else{//左括号压如栈中
			stack = append(stack,s[i])
		}
	}
	return len(stack)==0
	
}

func main(){
	fmt.Println(isValid("{([{])}{}()"))
}