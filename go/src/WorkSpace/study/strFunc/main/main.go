//系统函数 - 字符串的相关函数
package main

import "fmt"
func main(){
	str := "golang 你好"
	r := []rune(str)
	for i := 0;i < len(r) ;i++{
		fmt.Printf("%c \n",r[i])
	}
}