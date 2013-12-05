package main

import "fmt"
import "net"

var pre_def_dns = [] string{
	"seed.bitcoin.sipa.be", 
	"dnsseed.bluematt.me", 
	"dnsseed.bitcoin.dashjr.org", 
	"bitseed.xf2.org",
}

func pre_resolve_addr() {
	for _, element := range pre_def_dns {
		addr, e := net.LookupIP(element);
		fmt.Println("Total :", len(addr));
		fmt.Println("dns : ", element);
		fmt.Println(addr, e);
	}
}

func main() {
	pre_resolve_addr();
}
