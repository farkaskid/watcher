package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"watcher/icmp"
	"watcher/tcp"
)

func handleProto(proto string, bytes []byte) string {
	switch proto {
	case "tcp":
		packet := tcp.NewPacket(bytes)
		return packet.String()
	case "icmp":
		packet := icmp.NewPacket(bytes)
		return packet.String()
	}

	return proto + " packet"
}

func main() {
	protocol := flag.String("proto", "tcp", "the protocol to listen to.")
	flag.Parse()

	conn, err := net.ListenIP("ip4:"+*protocol, new(net.IPAddr))
	if err != nil {
		log.Panicln(err)
	}

	buf := make([]byte, 2048)
	for {
		count, ip, err := conn.ReadFrom(buf)
		if err != nil {
			break
		}

		domains, err := net.LookupAddr(ip.String())
		packetDes := handleProto(*protocol, buf)
		if err != nil {
			fmt.Println(packetDes, "from IP:", ip, "with", count, "bytes")
			continue
		}

		fmt.Println(packetDes, "from domain:", domains[0], "with", count, "bytes")
	}
}
