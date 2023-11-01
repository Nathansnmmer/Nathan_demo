package main

import (
	"fmt"
	"time"
)

// time demo
func main() {
	now := time.Now() //时间对象
	/*fmt.Println(now)
	year := now.Year()
	moth := now.Month()
	day := now.Day()
	hour := now.Hour()
	minit := now.Minute()
	sceond := now.Second()
	fmt.Println(year, moth, day, hour, minit, sceond)
	//时间戳: 从1970.1.1到现在的一个秒数
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1, timestamp2)
	//将一个时间戳转换一个具体的时间格式
	//1695475563+3600=1695479163
	t := time.Unix(1695479163, 0)
	fmt.Println(t)
	n := 5                                     //n是int类型不能直接*time second
	time.Sleep(time.Duration(n) * time.Second) //这里用time duration 强制转换n为duration类型
	fmt.Println("over")*/
	//加一小时的操作
	fmt.Println(now)
	t2 := now.Add(time.Hour)
	fmt.Println(t2)
	//两个时间相减sub
	fmt.Println(t2.Sub(now))
	//定时器
	//for tmp := range time.Tick(time.Second) {
	//	fmt.Println(tmp)
	//}

	//时间格式化 		y    m   d   H  M  S//年月日 时分秒
	//			  2006   01  02  15 04 05
	t3 := now.Format("2006-01-02")
	t4 := now.Format("2006-01-02 15：04：05")
	t5 := now.Format("2006-01-02 03：04：05 PM")
	t6 := now.Format("2006-01-02 03：04：05.000 PM")
	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println(t5)
	fmt.Println(t6)

	//解析字符串类型的时间
	times := "2023/09/25 20:00:00"
	//1 拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//根据时区解析一个字符串时间
	timeobj2, err := time.ParseInLocation("2006/01/02 15:04:05", times, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeobj2)
	timeobj, err := time.Parse("2006/01/02 15:04:05", times)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeobj)
}
