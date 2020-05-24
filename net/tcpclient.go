package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "192.168.1.23:4001")
	frame := make([]byte, 13, 13)

	for {
		for x := 0; x < 13; {
			length, err := conn.Read(frame[x:]) // 有数据则读，没有则阻塞
			if length == 0 {
				log.Println("读取长度为零??????")
				continue
			}
			if err != nil {
				log.Printf("长度: %d  数据: %X\n", length, frame)
				log.Println("读取出错: ", err)
				return
			}
			x += length
		}

		var ID int32

		帧格式 := ""
		if frame[0]&0x80 == 0 {
			帧格式 = "标准帧"
			ID = int32(frame[3]&0x07)<<8 | int32(frame[4])
		} else {
			帧格式 = "扩展帧"
			ID = int32(frame[1]&0x3F)<<24 | int32(frame[2]<<16) | int32(frame[3]<<8) | int32(frame[4])
		}

		帧类型 := ""
		if frame[0]&0x40 == 0 {
			帧类型 = "数据帧"
		} else {
			帧类型 = "远程帧"
			log.Println("远程帧????????????????????????????????????????????????????????????????????????🐒")
		}

		dataLen := frame[0] & 0x0F
		// _ = dataLen
		// _ = ID
		// _ = 帧格式
		// _ = 帧类型

		fmt.Printf("ID: %.2X\t帧类型: %s\t帧格式: %s\t数据长度: %X\t数据: %X\n", ID, 帧类型, 帧格式, dataLen, frame[5:])
		frame = make([]byte, 13)
	}

}
