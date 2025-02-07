package main

import "fmt"

// 继承的载体是结构体
type person struct {
	Name string
	Age  int
	Sex  string
}

func (*person) SayHello() {
	fmt.Println("this is from person")
}

// 匿名结构体模拟继承
// 匿名结构体在这里的体现，就是person，person没有属性名，他是个类型
type student1 struct {
	person
	School string
}

// 常规结构体嵌套，用的很多
// 嵌套结构体的使用，就是一个属性的概念，Person在student2中就是一个属性而已
type student2 struct {
	Person *person
	School string
}

func main() {
	// 匿名结构体继承
	stu := &student1{
		School: "middle",
	}
	// 匿名结构体这里自动做了处理 stu.person.Name = stu.Name
	stu.Name = "leo"
	stu.Age = 18
	stu.SayHello()
	fmt.Println(stu)

	stu2 := &student2{
		School: "small",
	}
	person := &person{
		Name: "ben",
		Age:  19,
	}
	stu2.Person = person
	stu2.Person.SayHello()
	fmt.Println(stu2)
}
