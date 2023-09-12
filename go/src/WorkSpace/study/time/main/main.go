//日期和时间函数
package main

import (
	"fmt"
	"time"//获取日期时间的函数包
)

func main(){
	//获取日期时间需要调用time下的Now函数
	now := time.Now()
	//打印返回值 和返回值对应的数据类型
	fmt.Printf("%v ~~~~~ 对应的类型为：%T \n",now,now)//2023-09-12 11:02:31.3382151 +0800 CST m=+0.001528301 ~~~~~ 对应的类型为：time.Time
	//对应的类型为：time.Time 是一个结构体类型

	//获取单独的年 月 日 时 分 秒 需要调用time.Time下的方法
	fmt.Printf("年：%v \n",now.Year())
	fmt.Printf("月：%v \n",int(now.Month()))//now.Month 返回的时月份的英文 想要月份返回数字则可以直接将 now下的Month函数的返回值强行转成int类型
	fmt.Printf("日：%v \n",now.Day())
	fmt.Printf("时：%v \n",now.Hour())
	fmt.Printf("分：%v \n",now.Minute())
	fmt.Printf("秒：%v \n",now.Second())

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//格式化日期格式 Y-m-d H:m:s
	fmt.Printf("%d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())//fmt.Printf直接将内容打印到控制台上
	data := fmt.Sprintf("%d-%d-%d %d:%d:%d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())//fmt.Sprintf会将内容返回给一个变量中
	fmt.Println(data)
}