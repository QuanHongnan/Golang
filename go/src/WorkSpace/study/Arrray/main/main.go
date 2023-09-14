package main

import (
	"fmt"
)

func OldTest()(int, int){
	//给出五个学生的成绩 求出成绩的总和 平均数
	scorel := 95
	score2 := 91
	score3 := 85
	score4 := 25
	score5 := 53
	//求和
	sum := scorel+score2+score3+score4+score5
	//平均数
	avg := sum/5
	//输出
	return sum,avg
}

func ArrayTest()(int,int){
	//定义数组
	var scores [5]int
	//将成绩存入数组
	scores[0] = 53
	scores[1] = 95
	scores[2] = 91
	scores[3] = 85
	scores[4] = 25
	
	//求和
	sum := 0
	for _,v := range scores{
		sum += v
	}
	//求平均值
	avg := sum/len(scores)
	return sum,avg
}

func main(){
	sum,avg := OldTest()
	fmt.Printf("成绩的总和为：%v，成绩的平均分为：%v \n",sum,avg)
	sum1,avg1 := ArrayTest()
	fmt.Printf("成绩的总和为：%v，成绩的平均分为：%v \n",sum1,avg1)
}