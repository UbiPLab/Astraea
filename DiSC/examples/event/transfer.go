package main

import (
	"fmt"
	"net"
)

func transfer() {
	fmt.Println("hello")

	//监听
	listen, err := net.Listen("tcp", "127.0.0.1 : 8002")
	fmt.Println("listen 80012")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listen.Close() //最后程序结束后才执行listen.Close( )进行关闭

	//阻塞等待用户链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		} // 一有连接就发一个协程去处理
		go handleConn(conn, priKey, response_ch)
	}
}
