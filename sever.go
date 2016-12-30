/*
***********************************************************
// Describe : TCP socket 通信 服务器端
// date  : 2016.12.28
// Author :
*********************************************************
*/

package main

import (
	"fmt"
	"go-file/Go_Chat/mysql"
	"log"
	"net"
)

var conns [5]net.Conn

func main() {
	//启动服务器&连接数据库
	StartSever()
}

func StartSever() {
	fmt.Println("Starting the server ...")

	//创建listener 用来监听和接收客户端的请求
	listener, err := net.Listen("tcp", "localhost:50001")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	//连接数据库
	mysql.OpenMysql()

	//监听并接受来自客户端的连接
	for {
		//客户端请求时会产生net.Conn类型的连接变量
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		//保存连接信息
		for i := 0; i < 5; i++ {
			if conns[i] == nil {
				conns[i] = conn
				break
			} else {
				//给在线的人提示有人上线了
				onLineNotice := "Someone online!"
				SendAndSaveMsg(conns[i], onLineNotice)
			}
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	var sendConn net.Conn

	for {
		//将收到的消息放进buf
		buf := make([]byte, 2048)

		len, err := conn.Read((buf))
		if err != nil {
			fmt.Println("error Reading", err.Error())
			return
		}
		fmt.Printf("%v\n", string(buf[:len]))
		//将收到的消息转发给另一客户端
		for i := 0; i < 5; i++ {
			if conns[i] != nil {
				sendConn = conns[i]
				//发送和保存
				SendAndSaveMsg(sendConn, string(buf[:len]))
				continue
			}
		}

	}
}

func SendAndSaveMsg(conn net.Conn, sendData string) {

	//LogInfo(conn.RemoteAddr().String())
	_, err := conn.Write([]byte(sendData))

	if err != nil {
		fmt.Println("Send Filed", err.Error())
		return
	}

}

func LogInfo(loginfor string) {
	log.Println(loginfor)
}
