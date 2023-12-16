package main

import (
	"fmt"
	"os"
)

func writeToFile(packetDetails string) {
	file, err := os.OpenFile("packetlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString(packetDetails + "\n")
}
