package icmp

import (
	"encoding/binary"
	"fmt"
)

// Packet represent a ICMP packet.
type Packet struct {
	Type, Code uint8
	Checksum   uint16
	Payload    []byte
}

func (p Packet) String() string {
	return fmt.Sprintf("ICMP packet:[type: %d]", p.Type)
}

// NewPacket creates a new ICMP packet.
func NewPacket(bytes []byte) *Packet {
	packet := new(Packet)
	packet.Type = uint8(bytes[0])
	packet.Code = uint8(bytes[1])
	packet.Checksum = binary.BigEndian.Uint16(bytes[2:4])
	packet.Payload = bytes[8:]

	return packet
}
