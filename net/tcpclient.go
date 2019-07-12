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
	total := 8000 * 24 * 3600
	index := 0
	// create bars
	barOK := pb.New(total).Prefix("正常帧数量")
	barErr := pb.New(total).Prefix("错误帧数量")
	// start pool
	pool, err := pb.StartPool(barOK, barErr)
	if err != nil {
		panic(err)
	}
	// refresh info every second (default 200ms)
	barOK.SetRefreshRate(time.Second)
	barErr.SetRefreshRate(time.Second)

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
			log.Println("错误长度: ", len(message))
			log.Printf("错误内容: %X\n", message)
			log.Println("***************************************")
			barErr.Increment()
			continue
		}

		barOK.Increment()
	}

	fmt.Println(total, "帧数据接收完毕，总运行时间： ", time.Since(took))
	pool.Stop()
}
