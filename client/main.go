package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/hkdnet/gsocket/lib"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s message", os.Args[0])
		os.Exit(1)
	}
	message := os.Args[1]

	serverIP := "localhost"
	serverPort := "55555"
	myIP := "localhost"
	myPort := 55556

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	lib.DealError(err)

	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(myIP)
	myAddr.Port = myPort
	conn, err := net.DialTCP("tcp", myAddr, tcpAddr)
	lib.DealError(err)
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))

	readBuf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readlen, err := conn.Read(readBuf)
	lib.DealError(err)

	fmt.Println("server: " + string(readBuf[:readlen]))
}
