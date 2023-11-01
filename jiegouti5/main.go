package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性和JSON的序列化

// 如果go语言包中定义的标识符（函数名，结构体名等等）首字母大写就是对外可见的
// 如果一个结构体中的字段首字母大写，那么该字段对外是可见的
// 但是可以使用tag来改变字段的名字使用反引号  ` `
type student struct {
	Name string `json:"name"` //这就是一个tag
	Id   int
}

// student 的构造函数
func newStudent(id int, name string) student {
	return student{
		Id:   id,
		Name: name,
	}

}

type class struct {
	Title    string
	Students []student
}

func main() {
	//创建一个班级变量c1
	c1 := class{
		Title:    "火箭班",
		Students: make([]student, 0, 20),
	}
	//往班级c1中添加学生
	for i := 0; i < 10; i++ {
		temp := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, temp)
	}
	fmt.Printf("%#v\n", c1)
	//JSON序列化就是把go语言中的数据->json格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal faied,err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	//json反序列化 把json的字符串->变成go语言中的数据
	jaon := `{"Title":"火箭班","Students":[{"Name":"狗鸡虚","Id":0},{"Name":"stu01","Id":1},{"Name":"stu02","Id":2},{"Name":"stu03","Id":3},{"Name":"stu04","Id":4}]}`
	var a class
	err = json.Unmarshal([]byte(jaon), &a)
	if err != nil {
		fmt.Println("json unmarshal faied,err", err)
		return
	}
	fmt.Printf("%#v\n", a)
}
