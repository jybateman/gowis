package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
)

var host string
var port int
var version bool
func init() {
	flag.StringVar(&host, "h", "",  "-h HOST	explicitly query HOST")
	flag.IntVar(&port, "p", 43,  "-p PORT	use port number PORT")
	flag.BoolVar(&version, "v", false,  "Display version")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] QUERY:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	getHost()
	TCPRequest(strings.Join(flag.Args(), " "))
}
