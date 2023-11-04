package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//udp client demo

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed :err%v\n", err)
			return
		}
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("recv from failed :err%v\n", err)
			return
		}
		fmt.Printf("read from %v,mag :%v\n", addr, string(buf[:n]))
	}
}
