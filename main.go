package main

import (
	"fmt"
	"flag"
	"strings"
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
	// fmt.Println(flag.Args())
	// flag.Usage()
	args := strings.Join(flag.Args(), " ")
	switch GetType(args) {
	case IPV4:
		fmt.Println(getIPWhois(args, "data/ipv4_list"))
		break
	case AS:
		fmt.Println(getASWhois(args, "data/as_list"))
		break
	case AS32:
		fmt.Println(getASWhois(args, "data/as32_list"))
		break
	case DNS:
		fmt.Println(getDNSWhois(args))
		break
	}
}
