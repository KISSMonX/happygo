package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cheggaaa/pb"
)

func main() {

	took := time.Now()
	// connect to this socket
	conn, _ := net.Dial("tcp", "192.168.1.23:4001")

	// 每秒 8000 多帧， 24小时大约 7 亿帧
	total := 9000 * 24 * 3600
	index := 0
	errCount := 0
	// create bars
	barOK := pb.New(total)
	barOK.SetRefreshRate(time.Second)
	barOK.Prefix("接收正常帧数")
	// bar will format numbers as (B, KiB, MiB, ...)
	barOK.SetUnits(pb.U_BYTES)
	barOK.Start()

	for {
		if index++; index > total {
			break
		}

		message, err := bufio.NewReader(conn).ReadBytes(0x07)
		if err != nil {
			log.Println("读取出错: ", err)
			return
		}

		if len(message) != 13 {
			errCount++
			log.Printf("错误长度: %d, 内容： %X\n", len(message), message)
			log.Printf("错误帧数： %d  *****************\n", errCount)

			continue
		}

		barOK.Add(len(message))
	}

	barOK.Finish()

	fmt.Println(total, "帧数据接收完毕，总运行时间： ", time.Since(took))
}
