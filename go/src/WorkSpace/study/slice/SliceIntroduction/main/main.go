//切片 切数组的片段 
package main
import (
	"fmt"
)

func main(){
	//定义数组
	var arr = [6]int{1,2,3,4,5,6}
	//定义切片 var slice []int = arr[2:6] 或者如下
	slice := arr[2:6] //切片是从 arr[2:6) 包含arr下标2 但不包含arr下标6也就是切片内的内容是arr下标2-5
	//切片的值
	fmt.Printf("slice的值%v\n",slice)
	//切片的长度
	fmt.Printf("slice的长度%v\n",len(slice))
	//切片的容量 公式 切片容量+切片底层数组的容量-切片起始索引
	fmt.Printf("slice的容量%v\n",cap(slice)) //6-2=4
	//slice地址
	fmt.Printf("slice的地址%p\n",&slice[0]) 
	fmt.Printf("arr[2]的地址%p\n",&arr[2]) 
}	