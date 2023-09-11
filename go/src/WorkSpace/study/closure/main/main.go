package main
import "fmt"

type function func(int)(int)

func getSum()(function){
	fmt.Println("闭包函数")
	var sum int = 0
	return func (num1 int)int{
		sum = sum + num1
		return sum
	}
}

func main(){
	var ser = make([] int ,2,3)
	
	fmt.Println(ser[1])
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
}
