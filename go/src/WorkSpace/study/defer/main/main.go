package main 
import "fmt"

func getSum(num ...int){
	var sum int = 0
	defer fmt.Println("sum 为",sum)
	for _,value := range num{
		sum = sum + value
	}
	fmt.Println(sum) 
}

func main(){
	getSum(10,20)
}

