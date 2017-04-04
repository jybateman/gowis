package main

import (
	"flag"
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
	flag.Parse()
	getHost()
}
