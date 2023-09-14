package main

import "fmt"

func details(){
	var arr = [3]int{1,2,3}
	copy(&arr)
	fmt.Printf("arr的数据类型为%T 值为：%v",arr,arr)
}

func copy(arr *[3]int){
	(*arr)[2] = 7 
}

func main(){
	details()
}