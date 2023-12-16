package main

import (
	"fmt"
	"net/http"

	"github.com/google/gopacket"
)

func sniffHTTP(packet gopacket.Packet) {
	applicationLayer := packet.ApplicationLayer()

	if applicationLayer != nil {
		payload := applicationLayer.Payload()
		if httpReq, err := http.ReadRequest(http.NoBody); err == nil {
			fmt.Printf("HTTP Request:\n%s %s %s\n", httpReq.Method, httpReq.URL, httpReq.Proto)
		}
		fmt.Printf("Payload (first 64 bytes): %x\n", payload[:min(len(payload), 64)])
	}
}
