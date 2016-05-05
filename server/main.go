package main

import (
	"fmt"
	"github.com/hkdnet/gsocket/lib"
	"net"
	"time"
)

func main() {
	service := ":55555"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	lib.DealError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	lib.DealError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	messageLen, err := conn.Read(messageBuf)
	lib.DealError(err)
	message := string(messageBuf[:messageLen])
	message = message + " too!"

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))
}
