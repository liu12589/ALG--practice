package main

import (
	"../输入输出总结/codec"
	"bufio"
	"io"
	"log"
	"net"
)

// 监听端口
// 接收客户端请求建立连接
// 创建 goroutine 处理连接
func pes(conn net.Conn) {
	// 处理完成后关闭连接
	defer conn.Close()
	// 读取连接中的内容
	read := bufio.NewReader(conn)
	for true {
		message, err := codec.Decode(read)
		if err != nil {
			if err == io.EOF {
				//如果报错为 EOF 则说明该连接中数据已读完
				log.Println("客户端连接关闭", conn.RemoteAddr().String())
				return
			}
			// 否则读取失败
			log.Println("read error", err.Error())
			return
		}
		log.Printf("客户端：%s 发送的消息为: %s", conn.RemoteAddr().String(), message)
		//向连接中写入 ok
		conn.Write([]byte("ok"))
	}
}

func main() {
	//监听进程
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("listen, err", err.Error())
		return
	}
	for true {
		//如果监听到有连接请求建立连接
		conn, err := lis.Accept()
		if err != nil {
			log.Println("conn err", err.Error())
		}
		// 开辟一个携程处理这个连接
		go pes(conn)
	}
}
