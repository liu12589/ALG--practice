package main

import (
	"../输入输出总结/codec"
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

/*
建立与服务端的连接
进行数据收发
关闭连接
*/

func main() {

	// 与改地址的tcp建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("conn err", err.Error())
		return
	}
	// 发送完请求后关闭连接
	defer conn.Close()
	data := bufio.NewReader(os.Stdin)
	for true {
		input, _ := data.ReadString('\n')
		input = strings.Trim(input, "\r\n")

		if input == "q" || input == "Q" {
			log.Println("退出链接")
			return
		}
		massage, _ := codec.Encode(input)
		conn.Write(massage)

		var p [512]byte
		//接收服务端发送来的请求
		n, err := conn.Read(p[:])
		if err != nil {
			if err == io.EOF {
				log.Println("服务端关闭链接", conn.RemoteAddr().String())
				return
			}
			log.Println("读取失败", err.Error())
			return
		}
		log.Println("服务端发送消息为", string(p[:n]))
	}
}

func test() (result []byte) {
	for i := 0; i < 50; i++ {
		result = append(result, byte(i))
	}
	return result
}
