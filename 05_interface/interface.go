package main

import "fmt"

// 定义一个接口
type Usber interface {
	getName() string
}

// 用手机去实现接口
type phone struct {
	Name string
}

func (p *phone) getName() string {
	return p.Name
}

// 用电脑去实现接口
type computer struct {
	Name string
}

func (c *computer) getName() string {
	return c.Name
}

// 封装一个公共方法，接口usber接口的各种实现
// 这个其实就是各种动物叫声的方法
// 特性：类型转换
func transData(u Usber) string {
	name := u.getName()
	return name + "处理后"
}

func main() {
	p := &phone{Name: "iphone"}
	c := &computer{Name: "mac"}
	//phone的实例化对象赋值给了接口，getName打印出的就是phone的属性
	var u Usber
	u = p
	fmt.Println(u.getName())
	newName1 := transData(p)
	newName2 := transData(c)
	fmt.Println(newName1, newName2)

	// 类型断言
	var u2 Usber
	u2 = &phone{Name: "小米"}
	fmt.Println(u2.getName())
	// 把u2转换成phone
	v, ok := u2.(*phone)
	if !ok {
		fmt.Println("类型断言失败")
	}
	fmt.Println(v.Name)
}
