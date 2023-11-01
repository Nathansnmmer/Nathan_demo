package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readfile() {
	//打开文件
	fileobj, err := os.Open("./xxx.txt") //打开文件只读 这个路径是相对路径./是和执行程序同一个目录下
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileobj.Close() //关闭文件
	//读取文件内容
	var tmp = make([]byte, 128)
	n, err := fileobj.Read(tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("read %d bytes from file\n", n)
	fmt.Println(string(tmp))
} //**1

// 打开文件
func readall() { //循环读取文件里的内容 内容比较多的情况下
	//打开文件
	fileobj, err := os.Open("./xxx.txt") //打开文件只读不能写 这个路径是相对路径./是和执行程序同一个目录下
	if err != nil {
		fmt.Println(err)
		return //return是退出
	}

	defer fileobj.Close() //关闭文件
	for {
		var tmp = make([]byte, 128)
		n, err := fileobj.Read(tmp)
		if err == io.EOF { //EOF就是end of file
			fmt.Println(string(tmp[:n])) //把当前读了多少个字节数据打印出来然后退出 [:n]这个是切片一下从0到n做切片
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("read %d bytes from file\n", n)
		fmt.Println(string(tmp[:n]))
	}

} //**2
// 1 这俩**1和**2都是read方法来读文件

// 2 bufio来读取文件 buf就是带缓冲区的  bufio先从磁盘读到缓冲区在读到代码中
func readBybufio() {
	fileobj, err := os.Open("./xxx.txt") //打开文件只读 这个路径是相对路径./是和执行程序同一个目录下
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileobj.Close() //关闭文件
	reader := bufio.NewReader(fileobj)
	for {
		lin, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(lin)
			return
		}
		if err != nil {
			fmt.Printf("read file failed err %v\n", err)
			return
		}
		fmt.Print(lin)
	}
}

// 3  ioutil 中的readfile//ioutil这个包里面的函数被移动到其他的地方了
// 注意iotuil尽量不要来读大的文件
func readByos() {
	cont, err := os.ReadFile("./xxx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(cont))
}

// write file
// 1 os.openfile
func write() {
	fileobj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) //TRUNC是清空原有的重新写入
	if err != nil {
		fmt.Printf("open file failed err", err)
		return
	}
	defer fileobj.Close()
	str := "我叫狗鸡虚"
	fileobj.Write([]byte(str))      //一个byte的切片
	fileobj.WriteString("hello 鸡虚") //string
}

// 2 bufio.newreader
func writeBybufio() { //bufio 写入的时候先写入缓冲区在写入磁盘
	fileobj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) //TRUNC是清空原有的重新写入
	if err != nil {
		fmt.Printf("open file failed err", err)
		return
	}
	defer fileobj.Close()
	write := bufio.NewWriter(fileobj)
	write.WriteString("狗鸡虚") // 将内容写入到缓冲区   为什么bufio比较快是先把内容写入缓冲区等缓冲区到一定的阈值的时候才会调用操作系统的文件io的接口把文件写入到磁盘里
	write.Flush()            //然后将缓冲区的内容写入磁盘	这里接着上面    所以它减少了频分的与磁盘io交互的过程所以用它来写大文件的时候会提高操作效率
}

// 3 os.writefile
func writeos() {
	str := "我命由我不由天"
	err := os.WriteFile("./xxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed err", err)
		return
	}
}

func main() {
	//readall()
	//readBybufio()
	//readByos()
	//write()
	//writeBybufio()
	writeos()
}
