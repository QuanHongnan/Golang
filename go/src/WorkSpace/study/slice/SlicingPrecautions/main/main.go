package main 
import (
	"fmt"
)

func main(){
	//定义一个数组
	var intarr [6]int = [6]int{1,2,3,4,5,6} 
	//定义一个 切片
	var slice []int = intarr[:3]//1,2,3
	//在定义一个切片
	slice2 := slice[:2]//1，2
	fmt.Println(slice2)
	//将slice2的第一个元素更改上面的intarr数组和slice切片的对应元素同样会被更改
	slice2[0]=66
	fmt.Println(intarr)
	fmt.Println(slice)

	//切片可以一定意义上的动态生长需要使用append()函数
	//原slice2 内容为 1，2
	slice2 = append(slice2,3,4,5)//1，2，3，4，5
	//底层原理 
	//1.底层追加元素的时候不是对原数组进行扩容 二十定义了一个新数组复制老数组的切片到新数组再追加
	//2.slice2指向的数组是新数组
	//3.往往我们在使用追加的时候其实想要的效果是给元切片追加 那只需要将原切片重新赋值
	//4.底层的新数组不能直接操作/维护
	//5.切片也可以追加切片 
	slice3 := []int{99,22}
	slice3 = append(slice3,slice...)//三个点不能省略 让系统知道追加的时一个切片  
	fmt.Println(slice3)//[99 22 66 2 3]
}