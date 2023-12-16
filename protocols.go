package main

import (
	"fmt"

	"github.com/google/gopacket"
)

func extractProtocols(packet gopacket.Packet) {
	networkLayer := packet.NetworkLayer()
	transportLayer := packet.TransportLayer()

	if networkLayer != nil {
		fmt.Printf("Network Layer: %s, ", networkLayer.LayerType())
	}

	if transportLayer != nil {
		fmt.Printf("Transport Layer: %s\n", transportLayer.LayerType())
	} else {
		fmt.Println("No Transport Layer")
	}
}
