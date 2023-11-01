package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统
// 1添加学员信息
// 2编辑学员信息
// 3展示所有学员的信息

type Student struct {
	ID   int
	Name string
	Age  int
}

var students []Student

func main() {
	for {
		fmt.Println("学员信息管理系统")
		fmt.Println("1. 添加学员信息")
		fmt.Println("2. 编辑学员信息")
		fmt.Println("3. 展示所有学员信息")
		fmt.Println("4. 退出")
		fmt.Print("请选择操作: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			editStudent()
		case 3:
			showStudents()
		case 4:
			fmt.Println("退出程序")
			os.Exit(0)
		default:
			fmt.Println("无效的选项，请重新选择")
		}
	}
}

func addStudent() {
	var student Student
	fmt.Print("请输入学员ID: ")
	fmt.Scanln(&student.ID)
	fmt.Print("请输入学员姓名: ")
	fmt.Scanln(&student.Name)
	fmt.Print("请输入学员年龄: ")
	fmt.Scanln(&student.Age)

	students = append(students, student)
	fmt.Println("学员信息添加成功")
}

func editStudent() {
	fmt.Print("请输入要编辑的学员ID: ")
	var id int
	fmt.Scanln(&id)

	for i := range students {
		if students[i].ID == id {
			fmt.Print("请输入新的学员姓名: ")
			fmt.Scanln(&students[i].Name)
			fmt.Print("请输入新的学员年龄: ")
			fmt.Scanln(&students[i].Age)
			fmt.Println("学员信息编辑成功")
			return
		}
	}
	fmt.Println("未找到学员信息")
}

func showStudents() {
	fmt.Println("学员信息列表：")
	for _, student := range students {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d\n", student.ID, student.Name, student.Age)
	}
}
