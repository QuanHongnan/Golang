//映射的操作
package main
import (
	"fmt"
)

func main(){
	//声明三个映射
	//声明映射方式一
	var map1 map[int]string
	map1 = make(map[int]string,5)
	//增加更新操作 映射名[key] = 值 当key不存在于此映射中就是增加 存在就是更新
	map1[0]="张三"
	map1[1]="李四"
	map1[2]="王五"
	map1[3]="赵六"
	// fmt.Println(map1)
	//声明映射方式二
	map2 := make(map[int]int)
	//增加更新操作
	map2[0] = 33
	map2[1] = 22
	map2[2] = 25
	map2[3] = 27	
	fmt.Println(map2)
	//声明映射方式三
	map3 := map[int]string{
	0:"张",
	1:"李",
	2:"王",
	}
	//增加更新操作
	map3[3]="赵"
	fmt.Println(map3)
	map3[2]="刘"
	fmt.Println(map3)
	//删除操作
	delete(map3,1)
	fmt.Println(map3)
	//清空操作没有一个专门的方法一次删除 只能通过遍历key 使用delete()方法
	//或声明一个map = make(...),make个新的,让原来的成为垃圾被gc回收
	//查找操作
	value,flag := map3[2]
	fmt.Println(value,flag)
}