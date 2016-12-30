/*
***********************************************************
// Describe : TCP socket 通信 客户端
// date  : 2016.12.28
// Author :
*********************************************************
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func IsError(err error) {
	if err != nil {
		fmt.Printf("BIN : %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	//创建连接：
	server := "127.0.0.1:50001"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	IsError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	IsError(err)

	//登录功能完成后该部分不需要
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("who are you?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		go handleConnection(conn)
		_, err = conn.Write([]byte(trimmedClient + ": " + trimmedInput))
	}

}
func handleConnection(conn net.Conn) {

	for {
		//将收到的消息放进buf
		buf := make([]byte, 2048)
		len, err := conn.Read((buf))
		IsError(err)
		fmt.Printf("%v\n", string(buf[:len]))

	}

}
