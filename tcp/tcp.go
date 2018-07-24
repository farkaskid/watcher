package tcp

import (
	"encoding/binary"
	"fmt"
)

// Packet represents a TCP packet.
type Packet struct {
	SrcPort, DstPort uint16
	SeqNum, ActNum   uint32
	Payload          []byte
}

func (p Packet) String() string {
	return fmt.Sprintf("TCP packet:[Dst Port: %d]", p.DstPort)
}

// NewPacket create a new TCP packet from the given bytes.
func NewPacket(bytes []byte) *Packet {
	p := new(Packet)
	p.SrcPort = binary.BigEndian.Uint16(bytes[:2])
	p.DstPort = binary.BigEndian.Uint16(bytes[2:4])
	p.SeqNum = binary.BigEndian.Uint32(bytes[4:8])
	p.ActNum = binary.BigEndian.Uint32(bytes[8:12])
	p.Payload = bytes[24:]

	return p
}
