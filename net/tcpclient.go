package main

import (
	"bufio"
	"log"
	"net"
)

func main() {

	cnt := 0
	// connect to this socket
	conn, _ := net.Dial("tcp", "192.168.1.101:8886")
	for {

		message, err := bufio.NewReader(conn).ReadBytes(0x07)
		if err != nil {
			log.Println("读取出错: ", err)
			return
		}
		cnt++

		if len(message) != 13 {
			log.Println("错误长度: ", len(message))
			log.Printf("错误内容: %X\n", message)
			log.Println("***************************************")
			continue
		}

		if cnt%500 == 0 {
			log.Println("正常长度: ", len(message))
			log.Printf("正常内容: %X\n", message)
			log.Println("============================")
			log.Println("已接收帧数量: ", cnt)
		}
	}
}
