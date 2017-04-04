package main

import (
	"os"
	"net"
	"fmt"
	"flag"
	"bufio"
	"strings"
	"strconv"
)

const (
	IPV4 = iota
	IPV6
	AS
	AS32
	DNS
	INV
)

func getIPWhois(str, file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for l, err := r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		data := strings.Fields(l)
		_, ipnet, err := net.ParseCIDR(data[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return ""
		}
		if ipnet.Contains(net.ParseIP(str)) {
			return data[1]
		}
	}
	return ""
}

func getASWhois(str, file string) string {
	var min int
	var max int
	
	asn, err := strconv.Atoi(str[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}	
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for l, err := r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		data := strings.Fields(l)
		r := strings.Split(data[0], "-")
		min, _ = strconv.Atoi(r[0])
		max, _ = strconv.Atoi(r[1])
		if asn >= min && asn <= max {
			return data[1]
		}
	}
	return ""
}


func getDNSWhois(str string) string {
	f, err := os.Open("data/domain_list")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	fields := strings.Split(str, ".")
	for l, err := r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		if strings.HasPrefix(l, "."+fields[len(fields)-1]) {
			whois := strings.Fields(l)
			if len(whois) > 1 {
				return whois[1]
			}
		}
	}
	return ""
}

func GetType(str string) int {
	ip := net.ParseIP(str)
	if ip != nil {
		if ip.To4() != nil {
			return IPV4
		} else {
			return IPV6
		}
	} else if strings.Index(strings.ToLower(str), "as") == 0 {
		if len(str) <= 7 {
			return AS
		} else {
			return AS32
		}
	} else if len(strings.Split(str, ".")) > 1 {
		return DNS
	}
	return INV
}

func getHost() {
	if len(host) == 0 {
		args := strings.Join(flag.Args(), " ")
		switch GetType(args) {
		case IPV4:
			host = getIPWhois(args, "data/ipv4_list")
			break
		case AS:
			host = getASWhois(args, "data/as_list")
			break
		case AS32:
			host = getASWhois(args, "data/as32_list")
			break
		case DNS:
			host = getDNSWhois(args)
			break
		default:
			fmt.Fprintln(os.Stderr, "No whois server is known for this kind of object.")
		}
	}
}
