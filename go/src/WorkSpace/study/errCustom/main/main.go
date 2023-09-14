//自定义错误
package main

import (
	"fmt"
	"errors"
)

func test()(err error){
	n1:=10
	n2:=0
	if n2 == 0 {
		return errors.New("除数不可为零")
	}else{
		fmt.Println(n1/n2)
		return nil
	}	
}

func main(){
	err := test()
	if err != nil{
		fmt.Println("错误信息：",err)
		panic(err)
	}
	fmt.Println("程序继续执行")
}