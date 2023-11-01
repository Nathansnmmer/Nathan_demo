package main

import "fmt"

type Student struct {
	Id    int
	Name  string
	Class string
}

func newStudent(id int, name, class string) *Student {
	return &Student{
		Id:    id,
		Name:  name,
		Class: class,
	}
}

type Studentmang struct {
	Students []*Student
}

// 这个是newstumag的一个构造函数
func newstumag() *Studentmang {
	return &Studentmang{
		Students: make([]*Student, 0, 20),
	}
}

// 定义方法 针对studentmang这个结构体写的方法
// 1 添加
func (s *Studentmang) addStu(stu *Student) { //添加方法   接收者是*studentmang这个结构体   传入一个stu参数 这个参数是*student类型（就是结构体类型）
	s.Students = append(s.Students, stu) //在用append把stu追加刀studentmang这个结构体中
} /*	这个s.students是调用student这个切片*/

// 2 更改信息
func (s *Studentmang) changeStu(stu *Student) {
	for i, v := range s.Students {
		if stu.Id == v.Id {
			s.Students[i] = stu
			return
		}
	}
	fmt.Printf("输入的学生信息有误 系统中没有学号是：%d这个的学生\n", stu.Id)
}

// 3 展示
func (s *Studentmang) showStu() {
	for _, v := range s.Students {
		fmt.Printf("学号:%d  姓名：%s  班级: %s\n", v.Id, v.Name, v.Class)
	}
}
