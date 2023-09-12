//错误处理机制
package main

import (
	"fmt"
)

func errHandling(){
	//使用defer+recover来捕获错误：defer后加上匿名函数调用
	defer func(){
		//调用recover内置函数，可以捕获错误：
		err:=recover()
		//如果没有捕获错误，返回值为零值：nil
		if err != nil{
			fmt.Println("错误已经捕获")
			fmt.Println("err是：",err)
		}
	}()
	n1:=10
	n2:=0
	fmt.Println(n1/n2)
}

func main(){
	errHandling()
	fmt.Println("函数后向下执行")
}

