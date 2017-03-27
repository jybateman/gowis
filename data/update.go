 package main

import (
	"os"
	"fmt"
	"time"
	"bytes"
	"regexp"
	"strings"
	"net/http"
	"io/ioutil"

	"golang.org/x/net/html"
)

const (
	URL = "https://www.iana.org"
	DNS = "https://www.iana.org/domains/root/db"
	IPV4 = "https://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.csv"
	IPV6 = "https://www.iana.org/assignments/ipv6-unicast-address-assignments/ipv6-unicast-address-assignments.csv"
	AS = "https://www.iana.org/assignments/as-numbers/as-numbers-1.csv"
	AS32 = "https://www.iana.org/assignments/as-numbers/as-numbers-2.csv"
)

func getWhois(lnk string) string {
	re := regexp.MustCompile(`\s+\S*$`)
	resp, err := http.Get(lnk)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	pg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	off := bytes.Index(pg, []byte("WHOIS Server:"))
	if off < 0 {
		return ""
	}
	end := bytes.Index(pg[off:], []byte("\n"))
	if end < 0 {
		return ""
	}
	return strings.Trim(string(re.Find(pg[off:off+end])), " ")
}

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

func getLinks(n *html.Node, lst *[]string) {
	if n.Data == "a" {
		// fmt.Println(n.Attr[0].Val)
		// fmt.Println(n.FirstChild.Data)
		*lst = append(*lst, n.Attr[0].Val, n.FirstChild.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getLinks(c, lst)
	}
}

func getDomain() {
	var lst []string

	f, err := os.Create("domain_list")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	resp, err := http.Get(DNS)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := findTable(doc)
	if n == nil {
		return
	}
	getLinks(n, &lst)
	for i := 0; i < len(lst); i += 2 {
		f.WriteString(lst[i+1]+" "+getWhois(URL+lst[i])+"\n")
		time.Sleep(time.Second*5)
	}
}

func main() {
	getDomain()
	// getWhois("https://www.iana.org/domains/root/db/com.html")
	// getWhois("https://www.iana.org/domains/root/db/abb.html")
}
