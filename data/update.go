 package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

const (
	DNS = "https://www.iana.org/domains/root/db"
	IPV4 = "https://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.csv"
	IPV6 = "https://www.iana.org/assignments/ipv6-unicast-address-assignments/ipv6-unicast-address-assignments.csv"
	AS = "https://www.iana.org/assignments/as-numbers/as-numbers-1.csv"
	AS32 = "https://www.iana.org/assignments/as-numbers/as-numbers-2.csv"
)

func findTable(n *html.Node) *html.Node {
	if n.Data == "table" {
		for _, tbl := range n.Attr {
			if tbl.Key == "id" && tbl.Val == "tld-table" {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t := findTable(c)
		if t != nil {
			return t
		}
	}
	return nil
}

func getLinks(n *html.Node, lvl int) {
	if n.Data == "a" {
		for _, a := range n.Attr {
			fmt.Println(a)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fmt.Println(c)
		}
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getLinks(c, lvl+1)
	}
}

func getDomain() {
	resp, _ := http.Get(DNS)
	doc, _ := html.Parse(resp.Body)
	
	n := findTable(doc)
	getLinks(n, 0)
	resp.Body.Close()
}

func main() {
	getDomain()
}
