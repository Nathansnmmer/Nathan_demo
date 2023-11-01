package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统
// 1添加学员信息
// 2编辑学员信息
// 3展示所有学员的信息

// 第一步先打印出系统界面
func showmeau() {
	fmt.Println("欢迎进入学员信息管理系统")
	fmt.Println("1 添加学员")
	fmt.Println("2 编辑学员信息")
	fmt.Println("3 展示所有学院的信息")
	fmt.Println("4 退出")
}

func jianpan() *Student { //定义了一个函数返回值是一个student结构体类型的值
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入")
	fmt.Print("请输入学生的学号: ")
	fmt.Scanf("%d\n", &id) //用户键入一个id取地址来更改id的初始值
	fmt.Print("请输入学生的姓名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生的班级: ")
	fmt.Scanf("%s\n", &class)
	sd := newStudent(id, name, class) //调用student结构体的构造函数来让   用户键入的三个值传入函数
	fmt.Println("信息更新成功")
	return sd

}
func main() {
	showmeau()        //调用第一步这个打印系统界面的函数
	sm := newstumag() //sm变量接受这个newstumg函数的返回值  此时sm相当于一个切片
	for {
		//这个是用scanf函数来使得界面可以输入一个数字
		var num int
		fmt.Print("请输入要操作的序号：")
		fmt.Scanf("%d\n", &num) //键入数字并且用&取地址来更改num的值
		fmt.Println("用户输入的是", num)
		switch num {
		case 1: //输入1就是添加
			sq := jianpan() //调用上面的jianpan函数让一个变量接受这个函数的返回值
			sm.addStu(sq)

		case 2: //输入2编辑学员
			sq := jianpan()
			sm.changeStu(sq)

		case 3: //输入3展示学员
			sm.showStu()

		case 4: //输入4退出

			os.Exit(0)

		}
	}
}
