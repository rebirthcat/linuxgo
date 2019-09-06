package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args)!=2 {
		fmt.Fprintf(os.Stderr,"Usage:%s hostname\n",os.Args[0])
		os.Exit(1)
	}

	hostname:=os.Args[1]

	addr,err:=net.ResolveIPAddr("ip",hostname)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Resoveld address is ",addr.String())
	os.Exit(0)

}
