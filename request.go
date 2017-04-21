package gowis

import (
	"io"
	"os"
	"net"
	"fmt"
	"bytes"
)

func TCPRequest(req string) {
	var buf bytes.Buffer
	
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(req+"\n"))
	io.Copy(&buf, conn)
	fmt.Print(buf.String())
}
