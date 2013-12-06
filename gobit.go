package main

import "os"
import "fmt"
import "net"
import "time"
import "math/rand"
import "io/ioutil"

var pre_def_dns = [] string{
	"seed.bitcoin.sipa.be", 
	"dnsseed.bluematt.me", 
	"dnsseed.bitcoin.dashjr.org", 
	"bitseed.xf2.org",
}

type msg_header struct {
	magic 		[4]byte;   // Magic value indicating message origin network, 
                          // and used to seek to next message when stream state is unknown
	command 	[12]byte; // ASCII string identifying the packet content, 
						  // NULL padded (non-NULL padding results in packet rejected)
	length		uint32;	  // Length of payload in number of bytes
	checksum 	uint32;	  // First 4 bytes of sha256(sha256(payload))
}

type version_msg struct {
	version 	int32;
	services 	uint64;
	timestamp	int64;
	addr_recv	[26]byte;
	addr_from	[26]byte;
	nonce		uint64;
}

func init_version_msg() {
	var header msg_header;

	header.magic[0] = 0xf9;
	header.magic[1] = 0xbe;
	header.magic[2] = 0xb4;
	header.magic[3] = 0xd9;

	cmd := []byte("version");
	copy(cmd[:12], header.command[:]);
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func pre_resolve_addr() {
	for _, element := range pre_def_dns {
		addr, err := net.LookupIP(element);
		if err != nil {
			fmt.Println("err:", err);
		}
		fmt.Println("------ Total :", len(addr), "------");
		fmt.Println("------ dns : ", element, "-------");
		for _, ele := range addr {
			fmt.Println(ele);
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func read_something() {
	var addr net.TCPAddr;
	addr.IP = net.IPv4(85, 214, 100, 177);
	addr.Port = 8333;

	conn, err := net.DialTCP("tcp", nil, &addr)
	checkError(err)
	fmt.Println("Connected...");

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println("result len:", len(result))
	fmt.Println(string(result))
}

func main() {
	pre_resolve_addr();
	read_something();
}
