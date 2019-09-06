package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var(
	remoteAddr string
	remotePort string
	localAddr  string
	localPort  string
)

func main() {

	flag.StringVar(&remoteAddr,"h","127.0.0.1","remoteAddr")
	flag.StringVar(&remotePort,"p","3333","remotePort")
	flag.StringVar(&localAddr,"lh","","localAddr")
	flag.StringVar(&localPort,"lp","","localPort")
	flag.Parse()

	var err error
	var tcpAddr *net.TCPAddr
	var localTcpAddr *net.TCPAddr
	var conn net.Conn

	var buf =make([]byte,1024)
	var temp []byte
	var inputReader=bufio.NewReader(os.Stdin)
	var n int

	tcpAddr,err=net.ResolveTCPAddr("tcp",remoteAddr+":"+remotePort)
	checkError(err)
	if localPort==""&&localAddr==""{
		localTcpAddr=nil
	}else {
		localTcpAddr,err=net.ResolveTCPAddr("tcp",localAddr+":"+localPort)
		checkError(err)
	}

	conn,err=net.DialTCP("tcp",localTcpAddr,tcpAddr)
	checkError(err)

	for true {

		temp,err=inputReader.ReadBytes('\n')
		checkError(err)
		_,err=conn.Write(temp[:len(temp)-1])
		if err!=nil {
			fmt.Fprintf(os.Stderr,"Fatal error:%s\n",err.Error())
			conn.Close()
			os.Exit(1)
		}

		n,err=conn.Read(buf)
		if err == io.EOF {
			fmt.Println("server closed")
			conn.Close()
			os.Exit(0)
		}else {
			fmt.Println(string(buf[0:n]))
		}
	}



}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error:%s\n",err.Error())
		os.Exit(1)
	}
}