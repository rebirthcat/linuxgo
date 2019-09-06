package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"Usage: %s hostname\n",os.Args[0])
		os.Exit(1)
	}

	hostname:=os.Args[1]

	addrs,err:=net.LookupAddr(hostname)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _,addr:=range addrs{
		fmt.Println(addr)
	}
	os.Exit(0)
}
