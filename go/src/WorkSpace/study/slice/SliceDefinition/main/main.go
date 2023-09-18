package main
import (
	"fmt"
)

func main(){
	//定义切片：make的三个参数 1.切片类型 2.切片长度 3.切片容量
	 slice := make([]int,4,20)
	 fmt.Println(slice)
	 fmt.Println("切片的长度为：",len(slice))
	 fmt.Println("撇片的容量为：",cap(slice))
	 slice[0] = 66
	 slice[1] = 88
	 fmt.Println(slice)

	slice2 := []int{1,4,7}
	fmt.Println(slice2)
	fmt.Println("切片的长度为：",len(slice2))
	fmt.Println("撇片的容量为：",cap(slice2))
}