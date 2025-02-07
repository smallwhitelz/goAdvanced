package main

import (
	"encoding/json"
	"fmt"
)

// 属性小写，序列化会没有结果
type student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	var s1 = &student{
		ID:     1,
		Gender: "male",
		Name:   "john",
		Sno:    "1001",
	}
	fmt.Println(s1)
	// json序列化
	strByte, _ := json.Marshal(s1)
	fmt.Println(string(strByte))
	// json反序列化
	var jsonStr = `{"ID":1,"Gender":"male","Name":"john","Sno":"1001"}`
	s2 := &student{}
	err := json.Unmarshal([]byte(jsonStr), s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
}
