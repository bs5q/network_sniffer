package main

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
)

func printPacketInfo(packet gopacket.Packet) {
	fmt.Printf("Time: %s, Length: %d\n", packet.Metadata().Timestamp.Format(time.Stamp), packet.Metadata().Length)
}
