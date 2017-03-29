package main

import (
	"fmt"
	"flag"
)

var host string
var port string
var version bool
func init() {
	flag.StringVar(&host, "h", "",  "-h HOST	explicitly query HOST")
	flag.StringVar(&port, "p", "",  "-p PORT	use port number PORT")
	flag.BoolVar(&version, "v", false,  "Display version")
}

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	flag.Usage()
}
