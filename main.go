package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	protocol := flag.String("proto", "tcp", "the protocol to listen to.")
	flag.Parse()

	conn, err := net.ListenIP("ip4:"+*protocol, new(net.IPAddr))
	if err != nil {
		log.Panicln(err)
	}

	protoname := strings.ToUpper(*protocol)
	buf := make([]byte, 2048)
	for {
		count, ip, err := conn.ReadFrom(buf)
		if err != nil {
			break
		}

		domains, err := net.LookupAddr(ip.String())
		if err != nil {
			fmt.Println(protoname, "packet from IP:", ip, "with", count, "bytes")
			continue
		}

		fmt.Println(protoname, "packet from domain:", domains, "with", count, "bytes")
	}
}
