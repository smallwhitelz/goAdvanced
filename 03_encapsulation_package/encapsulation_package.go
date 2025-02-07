package main

import (
	"fmt"
	encapsulation "gojson/03_encapsulation"
)

func main() {
	stu := encapsulation.NewStudent("zhangsan", 80)
	// score能打印，但是不能使用
	fmt.Println(stu)
	stu.SetScore(90)
	fmt.Println(stu.GetScore())
}
