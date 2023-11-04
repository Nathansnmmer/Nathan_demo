package main

import (
	"fmt"
	"net"
)

// UDP server demo

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("接收到的数据", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
