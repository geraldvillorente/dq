package main

import (
	"flag"
	"fmt"
	"net"

	c "github.com/fatih/color"
)

func main() {
	url := flag.String("url", "", "The URL to diagnose.")
	dnsQuery(*url)
}

/*
Get DNS information.
*/
func dnsQuery(url string) {
	p := c.New(c.FgGreen, c.Bold)
	fmt.Println()

	// Query IP address.
	p.Printf("Querying IP records...\n")
	iprecords, ipeErr := net.LookupIP(url)
	if ipeErr != nil {
		panic(ipeErr)
	}
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	fmt.Println()

	// Query CNAME records.
	p.Printf("Querying CNAME records...\n")
	cname, cnameErr := net.LookupCNAME(url)
	if cnameErr != nil {
		panic(cnameErr)
	}
	fmt.Printf("%s\n", cname)

	fmt.Println()

	// Query MX records.
	p.Printf("Querying MX records...\n")
	mxs, mxErr := net.LookupMX(url)
	if mxErr != nil {
		panic(mxErr)
	}
	for _, mx := range mxs {
		fmt.Printf("%s %v\n", mx.Host, mx.Pref)
	}

	fmt.Println()

	// Query NS records.
	// @Todo disabling as the package is buggy for now.
	/*
		p.Printf("Querying NS records...\n")
		nss, nsErr := net.LookupNS("www.messenger.com")

		//spew.Dump(nsErr)

		if nsErr != nil {
			panic(nsErr)
		}

		if len(nss) == 0 {
			fmt.Printf("no record")
		}
		for _, ns := range nss {
			fmt.Printf("%s\n", ns.Host)
		}

		fmt.Println()
	*/

	// Query txt records.
	p.Printf("Querying TXT records...\n")
	txts, txtErr := net.LookupTXT(url)
	if txtErr != nil {
		panic(txtErr)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		fmt.Printf("%s\n", txt)
	}

	fmt.Println()
}
