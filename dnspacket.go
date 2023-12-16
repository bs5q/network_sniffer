package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func sniffDNS(packet gopacket.Packet) {
	dnsLayer := packet.Layer(layers.LayerTypeDNS)
	if dnsLayer != nil {
		dnsPacket, _ := dnsLayer.(*layers.DNS)
		fmt.Printf("DNS Request: %s\n", dnsPacket.Questions[0].Name)
	}
}
