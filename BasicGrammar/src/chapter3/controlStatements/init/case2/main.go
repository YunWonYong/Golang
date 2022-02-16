package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

var (
	flag bool
)

func init() {
	go func() {

		fmt.Println("서버 오픈")
		ln, err := net.Listen("tcp", ":8089")
		if err != nil {
			fmt.Println(err)
		}
		for {
			c, err := ln.Accept()
			if err != nil {
				fmt.Println("server error", err)
				continue
			}
			go handleServerConnection(c)
		}
	}()
}

func handleServerConnection(c net.Conn) {
	defer c.Close()
	msg := ""
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received", msg)
	flag = true
}

func client(no int, msg string) {
	fmt.Printf("%d번째 사용자\n 전송 메시지: %s\n", no, msg)
	c, err := net.Dial("tcp", "localhost:8089")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	go client(1, "안녕")
	go client(2, "안녕")
	go client(1, "밥먹었어?")
	go client(2, "응")
	for {
		if flag {
			fmt.Println("서버 종료")
			break
		}
	}
}
