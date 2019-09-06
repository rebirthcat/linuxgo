package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var (
	host string
	port string
)

func main() {
	flag.StringVar(&host, "h", "", "host")
	flag.StringVar(&port, "p", "3333", "port")

	flag.Parse()
	var l *net.TCPListener
	var err error
	var conn net.Conn

	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	checkError(err)

	l, err = net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for true {
		conn, err = l.Accept()
		fmt.Printf("client:%s connect\n", conn.RemoteAddr())
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	var buf = make([]byte, 1024)
	var err error
	var n int
	var readstr string
	for true {
		n, err = conn.Read(buf)
		conn.Close()
		if err == io.EOF {
			fmt.Printf("client:%s closed\n", conn.RemoteAddr().String())
			conn.Close()
			return
		} else {
			readstr = strings.ToUpper(string(buf[0:n]))
			_, err = conn.Write([]byte(readstr))
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s\n", err.Error())
		os.Exit(1)
	}
}
