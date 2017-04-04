package main

import (
	"os"
	"net"
	"fmt"
)

func TCPRequest(req string) {
	b := make([]byte, 2048)
	fmt.Printf("%s:%d\n", host, port)
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	c.Write([]byte(req+"\n"))
	c.Read(b)
	fmt.Println(string(b))
}
